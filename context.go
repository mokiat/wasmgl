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
	initConstants(context)
	return nil
}

func ActiveTexture(texture int) {
	context.Call("activeTexture", texture)
}

func AttachShader(program Program, shader Shader) {
	context.Call("attachShader", js.Value(program), js.Value(shader))
}

func BindBuffer(target int, buffer Buffer) {
	context.Call("bindBuffer", target, js.Value(buffer))
}

func BindTexture(target int, texture Texture) {
	context.Call("bindTexture", target, js.Value(texture))
}

func BindVertexArray(array VertexArray) {
	context.Call("bindVertexArray", js.Value(array))
}

func Clear(mask int) {
	context.Call("clear", mask)
}

func ClearColor(r, g, b, a float32) {
	context.Call("clearColor", r, g, b, a)
}

func CompileShader(shader Shader) {
	context.Call("compileShader", js.Value(shader))
}

func CreateBuffer() Buffer {
	return Buffer(context.Call("createBuffer"))
}

func CreateProgram() Program {
	return Program(context.Call("createProgram"))
}

func CreateShader(shaderType int) Shader {
	return Shader(context.Call("createShader", shaderType))
}

func CreateTexture() Texture {
	return Texture(context.Call("createTexture"))
}

func CreateVertexArray() VertexArray {
	return VertexArray(context.Call("createVertexArray"))
}

func DeleteTexture(texture Texture) {
	context.Call("deleteTexture", js.Value(texture))
}

func DeleteVertexArray(array VertexArray) {
	context.Call("deleteVertexArray", js.Value(array))
}

func DepthFunc(fn int) {
	context.Call("depthFunc", fn)
}

func DepthMask(mask bool) {
	context.Call("depthMask", mask)
}

func Disable(cap int) {
	context.Call("disable", cap)
}

func DrawElements(mode, count, dtype, offset int) {
	context.Call("drawElements", mode, count, dtype, offset)
}

func Enable(cap int) {
	context.Call("enable", cap)
}

func GenerateMipmap(target int) {
	context.Call("generateMipmap", target)
}

func TexParameteri(target, pname, param int) {
	context.Call("texParameteri", target, pname, param)
}

func TexStorage2D(target, levels, internalFormat, width, height int) {
	context.Call("texStorage2D", target, levels, internalFormat, width, height)
}

func TexSubImage2D(target, level, xoffset, yoffset, width, height, format, dtype int, data []byte) {
	tData := newTypedSlice(data)
	defer tData.Release()
	context.Call("texSubImage2D", target, level, xoffset, yoffset, width, height, format, dtype, tData.JSUint8Array())
}

func UseProgram(program Program) {
	context.Call("useProgram", js.Value(program))
}

func VertexAttribPointer(index, size, dtype int, normalized bool, stride, offset int) {
	context.Call("vertexAttribPointer", index, size, dtype, normalized, stride, offset)
}

func Viewport(x, y, width, height int) {
	context.Call("viewport", x, y, width, height)
}
