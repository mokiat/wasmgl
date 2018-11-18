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
	gl.DepthTest = context.Get("DEPTH_TEST").Int()
	gl.CullFace = context.Get("CULL_FACE").Int()
	gl.ArrayBuffer = context.Get("ARRAY_BUFFER").Int()
	gl.ElementArrayBuffer = context.Get("ELEMENT_ARRAY_BUFFER").Int()
	gl.StaticDraw = context.Get("STATIC_DRAW").Int()
	gl.ColorBufferBit = context.Get("COLOR_BUFFER_BIT").Int()
	gl.DepthBufferBit = context.Get("DEPTH_BUFFER_BIT").Int()
	gl.VertexShader = context.Get("VERTEX_SHADER").Int()
	gl.FragmentShader = context.Get("FRAGMENT_SHADER").Int()
	gl.CompileStatus = context.Get("COMPILE_STATUS").Int()
	gl.LinkStatus = context.Get("LINK_STATUS").Int()
	gl.LEqual = context.Get("LEQUAL").Int()
	gl.Triangles = context.Get("TRIANGLES").Int()
	gl.Texture0 = context.Get("TEXTURE0").Int()
	gl.Texture1 = context.Get("TEXTURE1").Int()
	gl.Texture2 = context.Get("TEXTURE2").Int()
	gl.Texture2D = context.Get("TEXTURE_2D").Int()
	gl.TextureCubeMap = context.Get("TEXTURE_CUBE_MAP").Int()
	gl.TextureCubeMapPositiveX = context.Get("TEXTURE_CUBE_MAP_POSITIVE_X").Int()
	gl.TextureCubeMapNegativeX = context.Get("TEXTURE_CUBE_MAP_NEGATIVE_X").Int()
	gl.TextureCubeMapPositiveY = context.Get("TEXTURE_CUBE_MAP_POSITIVE_Y").Int()
	gl.TextureCubeMapNegativeY = context.Get("TEXTURE_CUBE_MAP_NEGATIVE_Y").Int()
	gl.TextureCubeMapPositiveZ = context.Get("TEXTURE_CUBE_MAP_POSITIVE_Z").Int()
	gl.TextureCubeMapNegativeZ = context.Get("TEXTURE_CUBE_MAP_NEGATIVE_Z").Int()
	gl.TextureWrapS = context.Get("TEXTURE_WRAP_S").Int()
	gl.TextureWrapT = context.Get("TEXTURE_WRAP_T").Int()
	gl.TextureWrapR = context.Get("TEXTURE_WRAP_R").Int()
	gl.TextureMinFilter = context.Get("TEXTURE_MIN_FILTER").Int()
	gl.TextureMagFilter = context.Get("TEXTURE_MAG_FILTER").Int()
	gl.ClampToEdge = context.Get("CLAMP_TO_EDGE").Int()
	gl.Linear = context.Get("LINEAR").Int()
	gl.RGBA = context.Get("RGBA").Int()
	gl.RGB = context.Get("RGB").Int()
	gl.Float = context.Get("FLOAT").Int()
	gl.UnsignedShort = context.Get("UNSIGNED_SHORT").Int()
	gl.UnsignedByte = context.Get("UNSIGNED_BYTE").Int()
	return gl, nil
}

// GL2 exposes WebGL2 API calls and constants.
type GL2 struct {
	context js.Value

	DepthTest               int
	CullFace                int
	ArrayBuffer             int
	ElementArrayBuffer      int
	StaticDraw              int
	ColorBufferBit          int
	DepthBufferBit          int
	VertexShader            int
	FragmentShader          int
	CompileStatus           int
	LinkStatus              int
	LEqual                  int
	Triangles               int
	Texture0                int
	Texture1                int
	Texture2                int
	Texture2D               int
	TextureCubeMap          int
	TextureCubeMapPositiveX int
	TextureCubeMapNegativeX int
	TextureCubeMapPositiveY int
	TextureCubeMapNegativeY int
	TextureCubeMapPositiveZ int
	TextureCubeMapNegativeZ int
	TextureWrapS            int
	TextureWrapT            int
	TextureWrapR            int
	TextureMinFilter        int
	TextureMagFilter        int
	ClampToEdge             int
	Linear                  int
	RGBA                    int
	RGB                     int
	Float                   int
	UnsignedShort           int
	UnsignedByte            int
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

func (gl *GL2) Enable(cap int) {
	gl.context.Call("enable", cap)
}

func (gl *GL2) DepthFunc(fn int) {
	gl.context.Call("depthFunc", fn)
}

func (gl *GL2) DepthMask(mask bool) {
	gl.context.Call("depthMask", mask)
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

func (gl *GL2) GetAttribLocation(program Program, name string) AttribLocation {
	location := gl.context.Call("getAttribLocation", program.id, name)
	return AttribLocation{location: location}
}

func (gl *GL2) GetUniformLocation(program Program, name string) UniformLocation {
	location := gl.context.Call("getUniformLocation", program.id, name)
	return UniformLocation{location: location}
}

func (gl *GL2) UseProgram(program Program) {
	gl.context.Call("useProgram", program.id)
}

func (gl *GL2) CreateTexture() Texture {
	id := gl.context.Call("createTexture")
	return Texture{id: id}
}

func (gl *GL2) BindTexture(target int, texture Texture) {
	gl.context.Call("bindTexture", target, texture.id)
}

func (gl *GL2) TexImage2D(target, level, internalFormat, width, height, border, format, dtype int, data []byte) {
	gl.context.Call("texImage2D", target, level, internalFormat, width, height, border, format, dtype, js.TypedArrayOf(data))
}

func (gl *GL2) ActiveTexture(texture int) {
	gl.context.Call("activeTexture", texture)
}

func (gl *GL2) TexParameteri(target, pname, param int) {
	gl.context.Call("texParameteri", target, pname, param)
}

func (gl *GL2) Uniform1i(location UniformLocation, x int) {
	gl.context.Call("uniform1i", location.location, x)
}

func (gl *GL2) UniformMatrix4fv(location UniformLocation, transpose bool, data []float32) {
	gl.context.Call("uniformMatrix4fv", location.location, transpose, js.TypedArrayOf(data))
}

func (gl *GL2) CreateVertexArray() VertexArray {
	id := gl.context.Call("createVertexArray")
	return VertexArray{id: id}
}

func (gl *GL2) BindVertexArray(array VertexArray) {
	gl.context.Call("bindVertexArray", array.id)
}

func (gl *GL2) DrawElements(mode, count, dtype, offset int) {
	gl.context.Call("drawElements", mode, count, dtype, offset)
}

func (gl *GL2) CreateBuffer() Buffer {
	id := gl.context.Call("createBuffer")
	return Buffer{id: id}
}

func (gl *GL2) BindBuffer(target int, buffer Buffer) {
	gl.context.Call("bindBuffer", target, buffer.id)
}

func (gl *GL2) BufferData(target int, data []byte, usage int) {
	// TODO: Handle nil data
	gl.context.Call("bufferData", target, js.TypedArrayOf(data), usage)
}

func (gl *GL2) EnableVertexAttribArray(index int) {
	gl.context.Call("enableVertexAttribArray", index)
}

func (gl *GL2) VertexAttribPointer(index, size, dtype int, normalized bool, stride, offset int) {
	gl.context.Call("vertexAttribPointer", index, size, dtype, normalized, stride, offset)
}

func (gl *GL2) Destroy() {
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

type AttribLocation struct {
	location js.Value
}

func (l AttribLocation) Valid() bool {
	// XXX: It is not actually documented what is returned on failure to allocate
	// so this is a best guess here
	return l.location.Int() != -1
}

type UniformLocation struct {
	location js.Value
}

func (l UniformLocation) Valid() bool {
	// XXX: It is not actually documented what is returned on failure to allocate
	// so this is a best guess here
	return l.location != js.Null()
}

type Texture struct {
	id js.Value
}

func (t Texture) Valid() bool {
	// XXX: It is not actually documented what is returned on failure to allocate
	// so this is a best guess here
	return t.id != js.Null()
}

type VertexArray struct {
	id js.Value
}

func (a VertexArray) Valid() bool {
	// XXX: It is not actually documented what is returned on failure to allocate
	// so this is a best guess here
	return a.id != js.Null()
}

type Buffer struct {
	id js.Value
}

func (b Buffer) Valid() bool {
	// XXX: It is not actually documented what is returned on failure to allocate
	// so this is a best guess here
	return b.id != js.Null()
}
