// +build js,wasm

package wasmgl

import (
	"fmt"
	"syscall/js"
)

func createGL(htmlCanvas js.Value) (*GL, error) {
	context := htmlCanvas.Call("getContext", "webgl")
	if context == js.Undefined() {
		return nil, fmt.Errorf("could not acquire webgl context")
	}
	return &GL{
		context:     context,
		glConstants: createGLConstants(context),
	}, nil
}

// GL exposes WebGL API calls and constants.
type GL struct {
	glConstants
	context js.Value
}

func (gl *GL) ActiveTexture(texture int) {
	gl.context.Call("activeTexture", texture)
}

func (gl *GL) AttachShader(program Program, shader Shader) {
	gl.context.Call("attachShader", program.id, shader.id)
}

func (gl *GL) BindBuffer(target int, buffer Buffer) {
	gl.context.Call("bindBuffer", target, buffer.id)
}

func (gl *GL) BindTexture(target int, texture Texture) {
	gl.context.Call("bindTexture", target, texture.id)
}

func (gl *GL) BufferData(target int, data []byte, usage int) {
	// TODO: Handle nil data
	typedArray := js.TypedArrayOf(data)
	defer typedArray.Release()
	gl.context.Call("bufferData", target, typedArray, usage)
}

func (gl *GL) Clear(mask int) {
	gl.context.Call("clear", mask)
}

func (gl *GL) ClearColor(r, g, b, a float64) {
	gl.context.Call("clearColor", r, g, b, a)
}

func (gl *GL) CompileShader(shader Shader) {
	gl.context.Call("compileShader", shader.id)
}

func (gl *GL) CreateBuffer() Buffer {
	id := gl.context.Call("createBuffer")
	return Buffer{id: id}
}

func (gl *GL) CreateProgram() Program {
	id := gl.context.Call("createProgram")
	return Program{id: id}
}

func (gl *GL) CreateShader(shaderType int) Shader {
	id := gl.context.Call("createShader", shaderType)
	return Shader{id: id}
}

func (gl *GL) CreateTexture() Texture {
	id := gl.context.Call("createTexture")
	return Texture{id: id}
}

func (gl *GL) DepthFunc(fn int) {
	gl.context.Call("depthFunc", fn)
}

func (gl *GL) DepthMask(mask bool) {
	gl.context.Call("depthMask", mask)
}

func (gl *GL) DrawElements(mode, count, dtype, offset int) {
	gl.context.Call("drawElements", mode, count, dtype, offset)
}

func (gl *GL) Enable(cap int) {
	gl.context.Call("enable", cap)
}

func (gl *GL) EnableVertexAttribArray(index int) {
	gl.context.Call("enableVertexAttribArray", index)
}

func (gl *GL) GetAttribLocation(program Program, name string) AttribLocation {
	location := gl.context.Call("getAttribLocation", program.id, name)
	return AttribLocation{location: location}
}

func (gl *GL) GetProgramParameter(program Program, pname int) Result {
	result := gl.context.Call("getProgramParameter", program.id, pname)
	return Result{value: result}
}

func (gl *GL) GetProgramInfoLog(program Program) string {
	result := gl.context.Call("getProgramInfoLog", program.id)
	return result.String()
}

func (gl *GL) GetShaderParameter(shader Shader, pname int) Result {
	result := gl.context.Call("getShaderParameter", shader.id, pname)
	return Result{value: result}
}

func (gl *GL) GetShaderInfoLog(shader Shader) string {
	result := gl.context.Call("getShaderInfoLog", shader.id)
	return result.String()
}

func (gl *GL) GetUniformLocation(program Program, name string) UniformLocation {
	location := gl.context.Call("getUniformLocation", program.id, name)
	return UniformLocation{location: location}
}

func (gl *GL) LinkProgram(program Program) {
	gl.context.Call("linkProgram", program.id)
}

func (gl *GL) ShaderSource(shader Shader, source string) {
	gl.context.Call("shaderSource", shader.id, source)
}

func (gl *GL) TexImage2D(target, level, internalFormat, width, height, border, format, dtype int, data []byte) {
	typedArray := js.TypedArrayOf(data)
	defer typedArray.Release()
	gl.context.Call("texImage2D", target, level, internalFormat, width, height, border, format, dtype, typedArray)
}

func (gl *GL) TexParameteri(target, pname, param int) {
	gl.context.Call("texParameteri", target, pname, param)
}

func (gl *GL) Uniform1i(location UniformLocation, x int) {
	gl.context.Call("uniform1i", location.location, x)
}

func (gl *GL) UniformMatrix4fv(location UniformLocation, transpose bool, data []float32) {
	typedArray := js.TypedArrayOf(data)
	defer typedArray.Release()
	gl.context.Call("uniformMatrix4fv", location.location, transpose, typedArray)
}

func (gl *GL) UseProgram(program Program) {
	gl.context.Call("useProgram", program.id)
}

func (gl *GL) VertexAttribPointer(index, size, dtype int, normalized bool, stride, offset int) {
	gl.context.Call("vertexAttribPointer", index, size, dtype, normalized, stride, offset)
}

func (gl *GL) Viewport(x, y, width, height int) {
	gl.context.Call("viewport", x, y, width, height)
}

func (gl *GL) Destroy() {
}

type AttribLocation struct {
	location js.Value
}

func (l AttribLocation) Valid() bool {
	return l.location.Int() != -1
}

type Buffer struct {
	id js.Value
}

func (b Buffer) Valid() bool {
	return b.id != js.Null()
}

type Program struct {
	id js.Value
}

func (p Program) Valid() bool {
	return p.id != js.Null()
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
	return s.id != js.Null()
}

type Texture struct {
	id js.Value
}

func (t Texture) Valid() bool {
	return t.id != js.Null()
}

type UniformLocation struct {
	location js.Value
}

func (l UniformLocation) Valid() bool {
	return l.location != js.Null()
}

func createGLConstants(context js.Value) glConstants {
	return glConstants{
		ARRAY_BUFFER:                context.Get("ARRAY_BUFFER").Int(),
		CLAMP_TO_EDGE:               context.Get("CLAMP_TO_EDGE").Int(),
		COLOR_BUFFER_BIT:            context.Get("COLOR_BUFFER_BIT").Int(),
		COMPILE_STATUS:              context.Get("COMPILE_STATUS").Int(),
		CULL_FACE:                   context.Get("CULL_FACE").Int(),
		DEPTH_BUFFER_BIT:            context.Get("DEPTH_BUFFER_BIT").Int(),
		DEPTH_TEST:                  context.Get("DEPTH_TEST").Int(),
		ELEMENT_ARRAY_BUFFER:        context.Get("ELEMENT_ARRAY_BUFFER").Int(),
		FLOAT:                       context.Get("FLOAT").Int(),
		FRAGMENT_SHADER:             context.Get("FRAGMENT_SHADER").Int(),
		LEQUAL:                      context.Get("LEQUAL").Int(),
		LINEAR:                      context.Get("LINEAR").Int(),
		LINK_STATUS:                 context.Get("LINK_STATUS").Int(),
		RGB:                         context.Get("RGB").Int(),
		RGBA:                        context.Get("RGBA").Int(),
		STATIC_DRAW:                 context.Get("STATIC_DRAW").Int(),
		TEXTURE0:                    context.Get("TEXTURE0").Int(),
		TEXTURE1:                    context.Get("TEXTURE1").Int(),
		TEXTURE2:                    context.Get("TEXTURE2").Int(),
		TEXTURE_2D:                  context.Get("TEXTURE_2D").Int(),
		TEXTURE_CUBE_MAP:            context.Get("TEXTURE_CUBE_MAP").Int(),
		TEXTURE_CUBE_MAP_POSITIVE_X: context.Get("TEXTURE_CUBE_MAP_POSITIVE_X").Int(),
		TEXTURE_CUBE_MAP_NEGATIVE_X: context.Get("TEXTURE_CUBE_MAP_NEGATIVE_X").Int(),
		TEXTURE_CUBE_MAP_POSITIVE_Y: context.Get("TEXTURE_CUBE_MAP_POSITIVE_Y").Int(),
		TEXTURE_CUBE_MAP_NEGATIVE_Y: context.Get("TEXTURE_CUBE_MAP_NEGATIVE_Y").Int(),
		TEXTURE_CUBE_MAP_POSITIVE_Z: context.Get("TEXTURE_CUBE_MAP_POSITIVE_Z").Int(),
		TEXTURE_CUBE_MAP_NEGATIVE_Z: context.Get("TEXTURE_CUBE_MAP_NEGATIVE_Z").Int(),
		TEXTURE_MAG_FILTER:          context.Get("TEXTURE_MAG_FILTER").Int(),
		TEXTURE_MIN_FILTER:          context.Get("TEXTURE_MIN_FILTER").Int(),
		TEXTURE_WRAP_S:              context.Get("TEXTURE_WRAP_S").Int(),
		TEXTURE_WRAP_T:              context.Get("TEXTURE_WRAP_T").Int(),
		TRIANGLES:                   context.Get("TRIANGLES").Int(),
		UNSIGNED_BYTE:               context.Get("UNSIGNED_BYTE").Int(),
		UNSIGNED_SHORT:              context.Get("UNSIGNED_SHORT").Int(),
		VERTEX_SHADER:               context.Get("VERTEX_SHADER").Int(),
	}
}

type glConstants struct {
	ARRAY_BUFFER                int
	CLAMP_TO_EDGE               int
	COLOR_BUFFER_BIT            int
	COMPILE_STATUS              int
	CULL_FACE                   int
	DEPTH_BUFFER_BIT            int
	DEPTH_TEST                  int
	ELEMENT_ARRAY_BUFFER        int
	FLOAT                       int
	FRAGMENT_SHADER             int
	LEQUAL                      int
	LINEAR                      int
	LINK_STATUS                 int
	RGB                         int
	RGBA                        int
	STATIC_DRAW                 int
	TEXTURE0                    int
	TEXTURE1                    int
	TEXTURE2                    int
	TEXTURE_2D                  int
	TEXTURE_CUBE_MAP            int
	TEXTURE_CUBE_MAP_POSITIVE_X int
	TEXTURE_CUBE_MAP_NEGATIVE_X int
	TEXTURE_CUBE_MAP_POSITIVE_Y int
	TEXTURE_CUBE_MAP_NEGATIVE_Y int
	TEXTURE_CUBE_MAP_POSITIVE_Z int
	TEXTURE_CUBE_MAP_NEGATIVE_Z int
	TEXTURE_MAG_FILTER          int
	TEXTURE_MIN_FILTER          int
	TEXTURE_WRAP_S              int
	TEXTURE_WRAP_T              int
	TRIANGLES                   int
	UNSIGNED_BYTE               int
	UNSIGNED_SHORT              int
	VERTEX_SHADER               int
}
