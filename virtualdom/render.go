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

func (v *VirtualDOM) Render(domChan chan string) {
	document := js.Global().Get("document")
	if v.root == "" {
		log.Fatal("virtual dom needs to be initialized")
	}
	root := document.Call("getElementById", v.root)
	for {
		dom := <-domChan
		root.Set("innerHTML", dom)
	}
}
