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

	gl := &GL2{
		context: context,
	}
	gl.ColorBufferBit = context.Get("COLOR_BUFFER_BIT").Int()
	gl.DepthBufferBit = context.Get("DEPTH_BUFFER_BIT").Int()
	return gl, nil
}

// GL2 exposes WebGL2 API calls and constants.
type GL2 struct {
	context js.Value

	ColorBufferBit int
	DepthBufferBit int
}

func (gl *GL2) Viewport(x, y, width, height int) {
	gl.context.Call("viewport", x, y, width, height)
}

func (gl *GL2) ClearColor(r, g, b, a float64) {
	gl.context.Call("clearColor", r, g, b, a)
}

func (gl *GL2) Clear(mask int) {
	gl.context.Call("clear", mask)
}
