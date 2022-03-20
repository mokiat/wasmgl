//go:build js && wasm

package wasmgl

import (
	"fmt"
	"syscall/js"
)

// References on WebGL1 and WebGL2 API:
// https://www.khronos.org/registry/webgl/specs/latest/1.0/
// https://www.khronos.org/registry/webgl/specs/latest/2.0/
// https://developer.mozilla.org/en-US/docs/Web/API/WebGLRenderingContext
// https://developer.mozilla.org/en-US/docs/Web/API/WebGL2RenderingContext

func GetExtension(name string) interface{} {
	result := context.Call("getExtension", name)
	if result.IsNull() {
		return nil
	}
	// TODO: In the future we could have specific Go types that are
	// returned depending on the extension name and that expose
	// the specific functions and constants of the extension.
	return true
}

func ReadPixels(x, y, width, height, format, dtype, offset int) {
	context.Call("readPixels", x, y, width, height, format, dtype, offset)
}

func ClientWaitSync(sync Sync, flags, timeout int) int {
	return context.Call("clientWaitSync", js.Value(sync), flags, timeout).Int()
}

func DeleteSync(sync Sync) {
	context.Call("deleteSync", js.Value(sync))
}

func FenceSync(condition, flags int) Sync {
	return Sync(context.Call("fenceSync", condition, flags))
}

func GetError() int {
	return context.Call("getError").Int()
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

func BlendFunc(sfactor, dfactor int) {
	context.Call("blendFunc", sfactor, dfactor)
}

func BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha int) {
	context.Call("blendFuncSeparate", srcRGB, dstRGB, srcAlpha, dstAlpha)
}

func BlendEquationSeparate(modeRGB, modeAlpha int) {
	context.Call("blendEquationSeparate", modeRGB, modeAlpha)
}

func Clear(mask int) {
	context.Call("clear", mask)
}

func ClearColor(r, g, b, a float32) {
	context.Call("clearColor", r, g, b, a)
}

func ClearDepth(depth float32) {
	context.Call("clearDepth", depth)
}

func ClearStencil(stencil int) {
	context.Call("clearStencil", stencil)
}

func ColorMask(r, g, b, a bool) {
	context.Call("colorMask", r, g, b, a)
}

func CompileShader(shader Shader) {
	context.Call("compileShader", js.Value(shader))
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

func CullFace(mode int) {
	context.Call("cullFace", mode)
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

func DisableVertexAttribArray(index int) {
	context.Call("disableVertexAttribArray", index)
}

func Enable(cap int) {
	context.Call("enable", cap)
}

func EnableVertexAttribArray(index int) {
	context.Call("enableVertexAttribArray", index)
}

func GenerateMipmap(target int) {
	context.Call("generateMipmap", target)
}

func GetAttribLocation(program Program, name string) AttribLocation {
	return AttribLocation(context.Call("getAttribLocation", js.Value(program), name))
}

func GetProgramInfoLog(program Program) string {
	return context.Call("getProgramInfoLog", js.Value(program)).String()
}

func GetProgramParameter(program Program, pname int) Result {
	return Result(context.Call("getProgramParameter", js.Value(program), pname))
}

func GetShaderInfoLog(shader Shader) string {
	return context.Call("getShaderInfoLog", js.Value(shader)).String()
}

func GetShaderParameter(shader Shader, pname int) Result {
	return Result(context.Call("getShaderParameter", js.Value(shader), pname))
}

func GetUniformLocation(program Program, name string) UniformLocation {
	return UniformLocation(context.Call("getUniformLocation", js.Value(program), name))
}

func LinkProgram(program Program) {
	context.Call("linkProgram", js.Value(program))
}

func ShaderSource(shader Shader, source string) {
	context.Call("shaderSource", js.Value(shader), source)
}

func TexImage2D(target, level, internalFormat, width, height, border, format, dtype int, data []byte) {
	tData := newTypedSlice(data)
	defer tData.Release()
	context.Call("texImage2D", target, level, internalFormat, width, height, border, format, dtype, tData.JSUint8Array())
}

func TexParameteri(target, pname, param int) {
	context.Call("texParameteri", target, pname, param)
}

func TexStorage2D(target, levels, internalFormat, width, height int) {
	context.Call("texStorage2D", target, levels, internalFormat, width, height)
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

// buffers

func CreateBuffer() Buffer {
	return Buffer(context.Call("createBuffer"))
}

func BufferData(target int, data []byte, usage int) {
	// TODO: Handle nil data
	tData := newTypedSlice(data)
	defer tData.Release()
	context.Call("bufferData", target, tData.JSUint8Array(), usage)
}

func BufferSubData(target, dstOffset int, data []byte) {
	tData := newTypedSlice(data)
	defer tData.Release()
	context.Call("bufferSubData", target, dstOffset, tData.JSUint8Array())
}

func GetBufferSubData(target, srcOffset int, data interface{}) {
	tData := newTypedSlice(data)
	defer tData.Release()
	uintArray := tData.JSUint8Array()
	context.Call("getBufferSubData", target, 0, uintArray)

	sliceData := sliceToByteSlice(data)
	js.CopyBytesToGo(sliceData, uintArray)
}

func DeleteBuffer(buffer Buffer) {
	context.Call("deleteBuffer", js.Value(buffer))
}

// drawing

func Scissor(x, y, width, height int) {
	context.Call("scissor", x, y, width, height)
}

func DrawArrays(mode, first, count int) {
	context.Call("drawArrays", mode, first, count)
}

func DrawElements(mode, count, dtype, offset int) {
	context.Call("drawElements", mode, count, dtype, offset)
}

// framebuffers

func CreateFramebuffer() Framebuffer {
	return Framebuffer(context.Call("createFramebuffer"))
}

func BindFramebuffer(target int, framebuffer Framebuffer) {
	context.Call("bindFramebuffer", target, js.Value(framebuffer))
}

func FramebufferTexture2D(target, attachment, texTarget int, texture Texture, level int) {
	context.Call("framebufferTexture2D", target, attachment, texTarget, js.Value(texture), level)
}

func CheckFramebufferStatus(target int) int {
	return context.Call("checkFramebufferStatus", target).Int()
}

func DrawBuffers(buffers []int) {
	// TODO: Figure out a more lightweight way to do this
	jsBuffers := make([]interface{}, len(buffers))
	for i, v := range buffers {
		jsBuffers[i] = v
	}
	context.Call("drawBuffers", jsBuffers)
}

func BlitFramebuffer(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter int) {
	context.Call("blitFramebuffer", srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter)
}

func DeleteFramebuffer(framebuffer Framebuffer) {
	context.Call("deleteFramebuffer", js.Value(framebuffer))
}

// shaders

func DeleteShader(shader Shader) {
	context.Call("deleteShader", js.Value(shader))
}

// programs

func DeleteProgram(program Program) {
	context.Call("deleteProgram", js.Value(program))
}

func DetachShader(program Program, shader Shader) {
	context.Call("detachShader", js.Value(program), js.Value(shader))
}

// stencil

func StencilFuncSeparate(face, fun, ref, mask int) {
	context.Call("stencilFuncSeparate", face, fun, ref, mask)
}

func StencilOpSeparate(face, fail, zfail, zpass int) {
	context.Call("stencilOpSeparate", face, fail, zfail, zpass)
}

// textures

func TexSubImage2D(target, level, xoffset, yoffset, width, height, format, dtype int, data []byte) {
	tData := newTypedSlice(data)
	defer tData.Release()

	switch dtype {
	case UNSIGNED_BYTE:
		buffer := tData.ArrayBuffer()
		uintArray := js.Global().Get("Uint8Array").New(buffer)
		context.Call("texSubImage2D", target, level, xoffset, yoffset, width, height, format, dtype, uintArray)
	case FLOAT:
		// TODO: Optimize this with staging area
		floatArray := js.Global().Get("Float32Array").New(tData.ArrayBuffer())
		context.Call("texSubImage2D", target, level, xoffset, yoffset, width, height, format, dtype, floatArray)
	default:
		panic(fmt.Errorf("unsupported dtype: %d", dtype))
	}
}

func TexSubImage3D(target, level, xoffset, yoffset, zoffset, width, height, depth, format, dtype int, data []byte) {
	tData := newTypedSlice(data)
	defer tData.Release()
	context.Call("texSubImage3D", target, level, xoffset, yoffset, zoffset, width, height, depth, format, dtype, tData.JSUint8Array())
}

func CopyTexSubImage2D(target, level, xoffset, yoffset, x, y, width, height int) {
	context.Call("copyTexSubImage2D", target, level, xoffset, yoffset, x, y, width, height)
}

// uniforms

func Uniform1f(location UniformLocation, x float32) {
	context.Call("uniform1f", js.Value(location), x)
}

func Uniform1i(location UniformLocation, x int) {
	context.Call("uniform1i", js.Value(location), x)
}

func Uniform3f(location UniformLocation, x, y, z float32) {
	context.Call("uniform3f", js.Value(location), x, y, z)
}

func Uniform4f(location UniformLocation, x, y, z, w float32) {
	context.Call("uniform4f", js.Value(location), x, y, z, w)
}

func UniformMatrix4fv(location UniformLocation, transpose bool, data []float32) {
	// TODO: Figure out a more lightweight way to do this
	jsData := make([]interface{}, len(data))
	for i, v := range data {
		jsData[i] = v
	}
	context.Call("uniformMatrix4fv", js.Value(location), transpose, jsData)
}
