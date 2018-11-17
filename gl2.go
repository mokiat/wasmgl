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
	gl.VertexShader = context.Get("VERTEX_SHADER").Int()
	gl.FragmentShader = context.Get("FRAGMENT_SHADER").Int()
	gl.CompileStatus = context.Get("COMPILE_STATUS").Int()
	gl.LinkStatus = context.Get("LINK_STATUS").Int()
	return gl, nil
}

// GL2 exposes WebGL2 API calls and constants.
type GL2 struct {
	context js.Value

	ColorBufferBit int
	DepthBufferBit int
	VertexShader   int
	FragmentShader int
	CompileStatus  int
	LinkStatus     int
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

func (gl *GL2) CreateShader(shaderType int) Shader {
	id := gl.context.Call("createShader", shaderType)
	return Shader{id: id}
}

func (gl *GL2) ShaderSource(shader Shader, source string) {
	gl.context.Call("shaderSource", shader.id, source)
}

func (gl *GL2) CompileShader(shader Shader) {
	gl.context.Call("compileShader", shader.id)
}

func (gl *GL2) GetShaderParameter(shader Shader, pname int) Result {
	result := gl.context.Call("getShaderParameter", shader.id, pname)
	return Result{value: result}
}

func (gl *GL2) GetShaderInfoLog(shader Shader) string {
	result := gl.context.Call("getShaderInfoLog", shader.id)
	return result.String()
}

func (gl *GL2) CreateProgram() Program {
	id := gl.context.Call("createProgram")
	return Program{id: id}
}

func (gl *GL2) AttachShader(program Program, shader Shader) {
	gl.context.Call("attachShader", program.id, shader.id)
}

func (gl *GL2) LinkProgram(program Program) {
	gl.context.Call("linkProgram", program.id)
}

func (gl *GL2) GetProgramParameter(program Program, pname int) Result {
	result := gl.context.Call("getProgramParameter", program.id, pname)
	return Result{value: result}
}

func (gl *GL2) GetProgramInfoLog(program Program) string {
	result := gl.context.Call("getProgramInfoLog", program.id)
	return result.String()
}

type Result struct {
	value js.Value
}

func (r Result) Bool() bool {
	return r.value.Bool()
}

func (r Result) Int() int {
	return r.value.Int()
}

type Shader struct {
	id js.Value
}

func (s Shader) Valid() bool {
	// XXX: It is not actually documented what is returned on failure to allocate
	// so this is a best guess here
	return s.id != js.Null()
}

type Program struct {
	id js.Value
}

func (p Program) Valid() bool {
	// XXX: It is not actually documented what is returned on failure to allocate
	// so this is a best guess here
	return p.id != js.Null()
}
