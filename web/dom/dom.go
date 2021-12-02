package dom

import (
	"strings"
	"syscall/js"
)

func GetElementById(id string) js.Value {
	return js.Global().Get("document").Call("getElementById", id)
}

func Clear(element js.Value) {
	if !element.IsUndefined() && !element.IsNull() {
		switch element.Get("nodeName").String() {
		case "INPUT", "TEXTAREA":
			element.Set("value", "")
		default:
			element.Set("textContent", "")
		}
	}
}

type domJs struct {
	JsValue js.Value
}

func NewEmptydomJs() *domJs {
	return &domJs{}
}

type domJsOption func(*domJs)

func NewdomJs(name string, opts ...domJsOption) js.Value {

	elementName := strings.ToLower(name)
	j := js.Global().Get("document").Call("createElement", elementName)

	ajs := &domJs{
		JsValue: j,
	}

	for _, opt := range opts {
		opt(ajs)
	}

	return ajs.JsValue
}

func Child(name string, opts ...domJsOption) domJsOption {
	return func(j *domJs) {
		elementName := strings.ToLower(name)
		c := js.Global().Get("document").Call("createElement", elementName)
		ajs := &domJs{
			JsValue: c,
		}
		for _, opt := range opts {
			opt(ajs)
		}
		j.JsValue.Call("appendChild", c)
	}
}

func ChildNamed(name string, ajsR *domJs, opts ...domJsOption) domJsOption {
	return func(j *domJs) {
		elementName := strings.ToLower(name)
		c := js.Global().Get("document").Call("createElement", elementName)

		ajs := &domJs{
			JsValue: c,
		}

		for _, opt := range opts {
			opt(ajs)
		}
		j.JsValue.Call("appendChild", c)
		ajsR.JsValue = c
	}
}

func Class(class string) domJsOption {
	return func(j *domJs) {
		j.JsValue.Call("setAttribute", "class", class)
	}
}

func Style(style string) domJsOption {
	return func(j *domJs) {
		j.JsValue.Call("setAttribute", "style", style)
	}
}

func Id(id string) domJsOption {
	return func(j *domJs) {
		j.JsValue.Call("setAttribute", "id", id)
	}
}

func Type(id string) domJsOption {
	return func(j *domJs) {
		j.JsValue.Call("setAttribute", "type", id)
	}
}

func Placeholder(id string) domJsOption {
	return func(j *domJs) {
		j.JsValue.Call("setAttribute", "placeholder", id)
	}
}

func Parent(parent js.Value) domJsOption {
	return func(j *domJs) {
		parent.Call("appendChild", j.JsValue)
	}
}

func Onclick(funcJs js.Func) domJsOption {
	return func(j *domJs) {
		j.JsValue.Set("onclick", funcJs)
	}
}

func Text(content string) domJsOption {
	return func(j *domJs) {
		elementType := j.JsValue.Get("nodeName").String()
		switch elementType {
		case "INPUT", "TEXTAREA":
			j.JsValue.Set("value", content)
		default:
			j.JsValue.Set("textContent", content)
		}

	}
}

func SetAttribute(attr, value string) domJsOption {
	return func(j *domJs) {
		j.JsValue.Call("setAttribute", attr, value)
	}
}

func PrintTeste(param string) domJsOption {
	return func(j *domJs) {
		println(param)
	}
}
