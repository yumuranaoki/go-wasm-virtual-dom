package virtualdom

import (
	"log"
	"syscall/js"
)

type VirtualDOM struct {
	root  string
	state interface{}
}

func Init() *VirtualDOM {
	v := &VirtualDOM{}
	return v
}

func (v *VirtualDOM) SetRoot(id string) {
	v.root = id
}

func (v *VirtualDOM) SetState(state interface{}) {
	v.state = state
}

func (v *VirtualDOM) RenderInitialState(html string) {
	document := js.Global().Get("document")
	if v.root == "" {
		log.Fatal("virtual dom needs to be initialized")
	}
	root := document.Call("getElementById", v.root)
	root.Set("innerHTML", html)
}

func (v *VirtualDOM) Render(html string) {
	c := make(chan bool)
	document := js.Global().Get("document")
	if v.root == "" {
		log.Fatal("virtual dom needs to be initialized")
	}
	root := document.Call("getElementById", v.root)
	root.Set("innerHTML", html)
	<-c
}
