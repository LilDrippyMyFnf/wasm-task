package main

import (
	"fmt"
	"strings"
	"syscall/js"

	. "github.com/jgbz/wasm-task/web/dom"
	"github.com/valyala/fastjson"
	fetch "marwan.io/wasm-fetch"
)

var waitw chan bool
var elements map[string]js.Value //mapa de armazenamento de elementos
var mapStatus map[string]string  //mapa de armazenamento de elementos

func init() {
	waitw = make(chan bool)
	elements = make(map[string]js.Value, 0)
	mapStatus = make(map[string]string)
	mapStatus["0"] = "Pending"
	mapStatus["1"] = "Doing"
	mapStatus["2"] = "Completed"
}

func main() {

	// Busco um elemento pelo id e atribui o valor à uma variável
	elements["divMain"] = GetElementById("main")

	// Limpa o conteudo de um elemento
	Clear(elements["divMain"])

	// time.Sleep(time.Millisecond * 900)
	// w := js.Global().Get("window")
	// w.Invoke(js.FuncOf(CreateElements))
	t := js.Global().Get("setTimeout")
	t.Invoke(js.FuncOf(CreateElements), 900)
	<-waitw //aguarda
}

// Função que cria os elementos html da pagina
func CreateElements(this js.Value, args []js.Value) interface{} {

	divMain := elements["divMain"]

	textArea := NewEmptydomJs()
	selectInput := NewEmptydomJs()
	divContent := NewEmptydomJs()

	NewdomJs("div",
		Parent(divMain),
		Class("container"),
		Child("section",
			Class("hero is-small has-background-primary mb-1"),
			Child("div",
				Class("hero-body"),
				Child("p",
					Text("Task"),
					Class("title has-text-white is-size-1 ml-1"),
				),
			),
		),
		Child("div",
			Class("field-body is-horizontal mt-4"),
			Child("div",
				Class("field"),
				Child("label",
					Text("Task"),
				),
				Child("div",
					Class("control"),
					ChildNamed("textarea",
						textArea,
						Class("input"),
						Type("textarea"),
						Placeholder("Groceries list"),
						SetAttribute("rows", "4"),
						SetAttribute("cols", "50"),
					),
				),
			),
			Child("div",
				Class("field"),
				Child("label",
					Text("Status"),
				),
				Child("div",
					Class("control"),
					Child("div",
						Class("select is-fullwidth"),
						ChildNamed("select",
							selectInput,
						),
					),
				),
			),
			Child("div",
				Class("field"),
				Child("div",
					Class("control"),
					Child("a",
						Class("button is-primary mt-5"),
						Text("Save"),
						Onclick(js.FuncOf(saveTask)),
					),
				),
			),
		),
		Child("nav",
			Class("level"),
			Child("p",
				Class("level-item has-text-centered"),
				Style("padding: 10px; border-bottom: 2px #00d1b2 solid;"),
			),
		),
		ChildNamed("div",
			divContent,
			Class("content"),
			Id("div_content"),
			Style("display: flex; flex-direction: column-reverse; gap: 10px;"),
		),
	)

	createStatusOptions(selectInput.JsValue)
	elements["textArea"] = textArea.JsValue
	elements["selectInput"] = selectInput.JsValue
	elements["divContent"] = divContent.JsValue
	getTask()
	return nil
}

func saveTask(this js.Value, args []js.Value) interface{} {
	task := elements["textArea"].Get("value").String()
	status := elements["selectInput"].Get("value").String()

	if len(task) < 1 {
		println("Invalid description")
		//Alert
		return nil
	}

	jsonStr := fmt.Sprintf(`{"description" : "%s", "status":"%s"}`, task, status)
	resp, err := fetch.Fetch("/v1/tasks", &fetch.Opts{
		Body:   strings.NewReader(jsonStr),
		Method: fetch.MethodPost,
	})

	if err != nil {
		println("erro")
	} else {
		if v, err := fastjson.ParseBytes(resp.Body); err != nil {
			println("erro")
		} else {
			id := string(v.GetStringBytes("id"))
			status := string(v.GetStringBytes("status"))

			if len(id) < 1 {
				//Alert
				return nil
			}
			createCard(id, task, status, elements["divContent"])
		}
	}
	return nil
}

func createStatusOptions(parent js.Value) {
	NewdomJs("option",
		Text("Pending"),
		SetAttribute("value", "0"),
		Parent(parent),
	)
	NewdomJs("option",
		Text("Doing"),
		SetAttribute("value", "1"),
		Parent(parent),
	)
	NewdomJs("option",
		Text("Completed"),
		SetAttribute("value", "2"),
		Parent(parent),
	)
}

func createCard(id, task, status string, parent js.Value) {

	buttons := NewEmptydomJs()
	NewdomJs("div",
		Parent(parent),
		Class("card w3-border w3-round-xlarge w3-bulma"),
		Id("task_"+id),
		Child("div",
			Class("card-content"),
			Style("display:flex;"),
			Child("div",
				Class("media-content"),
				Child("p",
					Class("title is-4 w3-text-white"),
					Text(task),
				),
				Child("p",
					Class("subtitle is-6 w3-text-white"),
					Id("subtitle_"+id),
					Text(mapStatus[status]),
				),
			),
			ChildNamed("div",
				buttons,
				Style("display:flex; gap: 5px;"),
			),
		),
	)

	switch status {
	case "0":
		buttonDoing(id, "1", buttons.JsValue)
		buttonDone(id, "2", buttons.JsValue)
	case "1":
		buttonDone(id, "2", buttons.JsValue)
	}

	btn := NewdomJs("a",
		Parent(buttons.JsValue),
		Class("button"),
		Child("i",
			Class("fas fa-trash-alt")),
	)

	btn.Set("idx", js.ValueOf(id))
	btn.Set("stX", js.ValueOf(status))
	btn.Set("onclick", js.FuncOf(deleteTask))

}

func buttonDoing(id, status string, parent js.Value) {

	btn := NewdomJs("a",
		Parent(parent),
		Class("button"),
		Child("i",
			Class("fas fa-play")),
	)

	btn.Set("idx", js.ValueOf(id))
	btn.Set("stX", js.ValueOf(status))
	btn.Set("onclick", js.FuncOf(updateTask))

}

func buttonDone(id, status string, parent js.Value) {

	btn := NewdomJs("a",
		Parent(parent),
		Class("button"),
		Child("i",
			Class("fas fa-check")),
	)

	btn.Set("idx", js.ValueOf(id))
	btn.Set("stX", js.ValueOf(status))
	btn.Set("onclick", js.FuncOf(updateTask))
}

func updateTask(this js.Value, args []js.Value) interface{} {
	id := this.Get("idx").String()
	status := this.Get("stX").String()
	btn := this

	jsonStr := fmt.Sprintf(`{"id" : "%s", "status":"%s"}`, id, status)

	resp, err := fetch.Fetch("/v1/tasks", &fetch.Opts{
		Body:   strings.NewReader(jsonStr),
		Method: fetch.MethodPatch,
	})

	if err != nil {
		println("erro")
	} else {

		if v, err := fastjson.ParseBytes(resp.Body); err != nil {
			println("erro")
		} else {
			id := string(v.GetStringBytes("id"))
			status := string(v.GetStringBytes("status"))

			if len(id) < 1 {
				//Alert
				return nil
			}

			element := GetElementById("subtitle_" + id)
			element.Set("innerHTML", mapStatus[status])
			btn.Call("remove")
		}
	}

	return nil
}

func deleteTask(this js.Value, args []js.Value) interface{} {
	id := this.Get("idx").String()

	jsonStr := fmt.Sprintf(`{"id" : "%s"}`, id)

	resp, err := fetch.Fetch("/v1/tasks", &fetch.Opts{
		Body:   strings.NewReader(jsonStr),
		Method: fetch.MethodDelete,
	})

	if err != nil {
		println("erro")
	} else {

		if v, err := fastjson.ParseBytes(resp.Body); err != nil {
			println("erro")
		} else {
			id := string(v.GetStringBytes("id"))

			if len(id) < 1 {
				//Alert
				return nil
			}

			GetElementById("task_" + id).Call("remove")
		}
	}

	return nil
}

func getTask() {

	resp, err := fetch.Fetch("/v1/tasks", &fetch.Opts{
		Method: fetch.MethodGet,
	})

	if err != nil {
		println("erro get")
	} else {
		var sc fastjson.Scanner

		sc.Init(string(resp.Body))
		for sc.Next() {
			arr := sc.Value().GetArray("tasks")
			for _, v := range arr {
				id := string(v.GetStringBytes("id"))
				task := string(v.GetStringBytes("description"))
				status := string(v.GetStringBytes("status"))
				createCard(id, task, status, elements["divContent"])
			}
		}
	}
}
