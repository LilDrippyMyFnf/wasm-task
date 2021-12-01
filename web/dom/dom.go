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

type adomJs struct {
	JsValue js.Value
}

func NewEmptyAdomJs() *adomJs {
	return &adomJs{}
}

type adomJsOption func(*adomJs)

func NewAdomJs(name string, opts ...adomJsOption) js.Value {

	elementName := strings.ToLower(name)
	j := js.Global().Get("document").Call("createElement", elementName)

	ajs := &adomJs{
		JsValue: j,
	}

	for _, opt := range opts {
		opt(ajs)
	}

	return ajs.JsValue
}

func Child(name string, opts ...adomJsOption) adomJsOption {
	return func(j *adomJs) {
		elementName := strings.ToLower(name)
		c := js.Global().Get("document").Call("createElement", elementName)
		ajs := &adomJs{
			JsValue: c,
		}
		for _, opt := range opts {
			opt(ajs)
		}
		j.JsValue.Call("appendChild", c)
	}
}

func Child2(name string, ajsR *adomJs, opts ...adomJsOption) adomJsOption {
	return func(j *adomJs) {
		elementName := strings.ToLower(name)
		c := js.Global().Get("document").Call("createElement", elementName)

		ajs := &adomJs{
			JsValue: c,
		}

		for _, opt := range opts {
			opt(ajs)
		}
		j.JsValue.Call("appendChild", c)
		ajsR.JsValue = c
	}
}

func Class(class string) adomJsOption {
	return func(j *adomJs) {
		j.JsValue.Call("setAttribute", "class", class)
	}
}

func Style(style string) adomJsOption {
	return func(j *adomJs) {
		j.JsValue.Call("setAttribute", "style", style)
	}
}

func Id(id string) adomJsOption {
	return func(j *adomJs) {
		j.JsValue.Call("setAttribute", "id", id)
	}
}

func Type(id string) adomJsOption {
	return func(j *adomJs) {
		j.JsValue.Call("setAttribute", "type", id)
	}
}

func Placeholder(id string) adomJsOption {
	return func(j *adomJs) {
		j.JsValue.Call("setAttribute", "placeholder", id)
	}
}

func Parent(parent js.Value) adomJsOption {
	return func(j *adomJs) {
		parent.Call("appendChild", j.JsValue)
	}
}

func Onclick(funcJs js.Func) adomJsOption {
	return func(j *adomJs) {
		j.JsValue.Set("onclick", funcJs)
	}
}

func Text(content string) adomJsOption {
	return func(j *adomJs) {
		elementType := j.JsValue.Get("nodeName").String()
		switch elementType {
		case "INPUT", "TEXTAREA":
			j.JsValue.Set("value", content)
		default:
			j.JsValue.Set("textContent", content)
		}

	}
}
