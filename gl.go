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
	gl.context.Call("bufferData", target, js.TypedArrayOf(data), usage)
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
	gl.context.Call("texImage2D", target, level, internalFormat, width, height, border, format, dtype, js.TypedArrayOf(data))
}

func (gl *GL) TexParameteri(target, pname, param int) {
	gl.context.Call("texParameteri", target, pname, param)
}

func (gl *GL) Uniform1i(location UniformLocation, x int) {
	gl.context.Call("uniform1i", location.location, x)
}

func (gl *GL) UniformMatrix4fv(location UniformLocation, transpose bool, data []float32) {
	gl.context.Call("uniformMatrix4fv", location.location, transpose, js.TypedArrayOf(data))
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
		ArrayBuffer:             context.Get("ARRAY_BUFFER").Int(),
		ClampToEdge:             context.Get("CLAMP_TO_EDGE").Int(),
		ColorBufferBit:          context.Get("COLOR_BUFFER_BIT").Int(),
		CompileStatus:           context.Get("COMPILE_STATUS").Int(),
		CullFace:                context.Get("CULL_FACE").Int(),
		DepthBufferBit:          context.Get("DEPTH_BUFFER_BIT").Int(),
		DepthTest:               context.Get("DEPTH_TEST").Int(),
		ElementArrayBuffer:      context.Get("ELEMENT_ARRAY_BUFFER").Int(),
		Float:                   context.Get("FLOAT").Int(),
		FragmentShader:          context.Get("FRAGMENT_SHADER").Int(),
		LEqual:                  context.Get("LEQUAL").Int(),
		Linear:                  context.Get("LINEAR").Int(),
		LinkStatus:              context.Get("LINK_STATUS").Int(),
		RGB:                     context.Get("RGB").Int(),
		RGBA:                    context.Get("RGBA").Int(),
		StaticDraw:              context.Get("STATIC_DRAW").Int(),
		Texture0:                context.Get("TEXTURE0").Int(),
		Texture1:                context.Get("TEXTURE1").Int(),
		Texture2:                context.Get("TEXTURE2").Int(),
		Texture2D:               context.Get("TEXTURE_2D").Int(),
		TextureCubeMap:          context.Get("TEXTURE_CUBE_MAP").Int(),
		TextureCubeMapPositiveX: context.Get("TEXTURE_CUBE_MAP_POSITIVE_X").Int(),
		TextureCubeMapNegativeX: context.Get("TEXTURE_CUBE_MAP_NEGATIVE_X").Int(),
		TextureCubeMapPositiveY: context.Get("TEXTURE_CUBE_MAP_POSITIVE_Y").Int(),
		TextureCubeMapNegativeY: context.Get("TEXTURE_CUBE_MAP_NEGATIVE_Y").Int(),
		TextureCubeMapPositiveZ: context.Get("TEXTURE_CUBE_MAP_POSITIVE_Z").Int(),
		TextureCubeMapNegativeZ: context.Get("TEXTURE_CUBE_MAP_NEGATIVE_Z").Int(),
		TextureMagFilter:        context.Get("TEXTURE_MAG_FILTER").Int(),
		TextureMinFilter:        context.Get("TEXTURE_MIN_FILTER").Int(),
		TextureWrapS:            context.Get("TEXTURE_WRAP_S").Int(),
		TextureWrapT:            context.Get("TEXTURE_WRAP_T").Int(),
		Triangles:               context.Get("TRIANGLES").Int(),
		UnsignedByte:            context.Get("UNSIGNED_BYTE").Int(),
		UnsignedShort:           context.Get("UNSIGNED_SHORT").Int(),
		VertexShader:            context.Get("VERTEX_SHADER").Int(),
	}
}

type glConstants struct {
	ArrayBuffer             int
	ClampToEdge             int
	ColorBufferBit          int
	CompileStatus           int
	CullFace                int
	DepthTest               int
	DepthBufferBit          int
	ElementArrayBuffer      int
	Float                   int
	FragmentShader          int
	LEqual                  int
	Linear                  int
	LinkStatus              int
	RGB                     int
	RGBA                    int
	StaticDraw              int
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
	TextureMagFilter        int
	TextureMinFilter        int
	TextureWrapS            int
	TextureWrapT            int
	Triangles               int
	UnsignedByte            int
	UnsignedShort           int
	VertexShader            int
}
