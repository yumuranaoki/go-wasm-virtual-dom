package virtualdom

import (
	"syscall/js"
	"log"
)

type VirtualDOM struct {
	root string
}

func (v VirtualDOM) setRoot(id string) {

}

func (v VirtualDOM) initialRendering(html string) {
	
}

func (v VirtualDOM) rendering() {
	document := js.Global().Get("document")
	if v.root == nil {
		log.Fatal("virtual dom needs to be initialized")
	}
	root := document.Call("getElementById", id)
}