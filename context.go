//go:build js && wasm

package wasmgl

import (
	"fmt"
	"syscall/js"
)

var context js.Value

// InitFromID initializes webgl context and bindings
// from the canvas that has the specified canvasID ID.
func InitFromID(canvasID string) error {
	htmlDocument := js.Global().Get("document")
	if htmlDocument.IsUndefined() {
		return fmt.Errorf("could not locate document element")
	}
	htmlCanvas := htmlDocument.Call("getElementById", canvasID)
	if htmlCanvas.IsNull() {
		return fmt.Errorf("could not locate canvas element with id %q", canvasID)
	}
	return InitFromCanvas(htmlCanvas)
}

// InitFromCanvas initializes webgl context and bindings
// from the specified htmlCanvas canvas element reference.
func InitFromCanvas(htmlCanvas js.Value) error {
	// TODO: Consider passing config options to getContext
	// so that depth buffer and other options can be specified
	// including performance hints
	context = htmlCanvas.Call("getContext", "webgl2")
	if context.IsNull() {
		return fmt.Errorf("could not acquire webgl2 context")
	}
	return nil
}
