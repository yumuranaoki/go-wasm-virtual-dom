package main

import (
	"strconv"
	"syscall/js"

	"github.com/yumuranaoki/go-wasm-virtual-dom/virtualdom"
)

type State struct {
	count int
}

var state = &State{
	count: 0,
}

func html(varibales string) string {
	return `<div><h1>Count</h1><p id="value-1">` + varibales + `</p><button onclick="increment()">+</button></div>`
}

func increment(this js.Value, values []js.Value) interface{} {
	state.count++
	virtualdom := virtualdom.Init()
	virtualdom.SetRoot("root")
	virtualdom.SetState(state)
	renderedHTML := html(strconv.Itoa(state.count))
	virtualdom.RenderInitialState(renderedHTML)
	return nil
}

func main() {
	virtualdom := virtualdom.Init()
	virtualdom.SetRoot("root")
	virtualdom.SetState(state)
	renderedHTML := html(strconv.Itoa(state.count))
	virtualdom.RenderInitialState(renderedHTML)
	js.Global().Set("increment", js.FuncOf(increment))
	// htmlをchanelで渡す
	virtualdom.Render(renderedHTML)
}
