package main

import (
	"syscall/js"
	"time"

	. "github.com/jgbz/wasm-task/web/dom"
)

var waitw chan bool
var elements map[string]js.Value //mapa de armazenamento de elementos
var mapStatus map[string]string  //mapa de armazenamento de elementos

func init() {
	waitw = make(chan bool)
	elements = make(map[string]js.Value, 0)
	mapStatus = make(map[string]string)
	mapStatus["0"] = "Pendente"
	mapStatus["1"] = "Fazendo"
	mapStatus["2"] = "Concluida"
}

func main() {

	// Busco um elemento pelo id e atribui o valor à uma variável
	elements["divMain"] = GetElementById("div_main")

	// Limpa o conteudo de um elemento
	Clear(divMain)

	time.Sleep(time.Millisecond * 900)
	js.Call(js.FuncOf(Create))
	<-waitw //aguarda
}

// Função que cria os elementos html da pagina
func Create(this js.Value, args []js.Value) interface{} {
	divMain := elements["divMain"]

	return nil
}
