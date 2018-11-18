// +build js,wasm

package wasmgl

import (
	"fmt"
	"syscall/js"
)

func createGL2(htmlCanvas js.Value) (*GL2, error) {
	context := htmlCanvas.Call("getContext", "webgl2")
	if context == js.Undefined() {
		return nil, fmt.Errorf("could not acquire webgl2 context")
	}
	return &GL2{
		GL: GL{
			glConstants: createGLConstants(context),
			context:     context,
		},
		gl2Constants: createGL2Constants(context),
	}, nil
}

// GL2 exposes WebGL2 API calls and constants.
type GL2 struct {
	GL
	gl2Constants
}

func (gl *GL2) BindVertexArray(array VertexArray) {
	gl.context.Call("bindVertexArray", array.id)
}

func (gl *GL2) CreateVertexArray() VertexArray {
	id := gl.context.Call("createVertexArray")
	return VertexArray{id: id}
}

func (gl *GL2) Destroy() {
}

type VertexArray struct {
	id js.Value
}

func (a VertexArray) Valid() bool {
	return a.id != js.Null()
}

func createGL2Constants(context js.Value) gl2Constants {
	return gl2Constants{
		TEXTURE_WRAP_R: context.Get("TEXTURE_WRAP_R").Int(),
	}
}

type gl2Constants struct {
	TEXTURE_WRAP_R int
}
