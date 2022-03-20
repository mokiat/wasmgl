//go:build js && wasm

package wasmgl

import (
	"fmt"
	"syscall/js"
)

var (
	// NOTE: We use references to JS functions and Invoke instead of
	// Call since the latter leads to strings being passed around
	// and TextDecoder being used on JS side.

	// WebGL1 functions
	// (https://www.khronos.org/registry/webgl/specs/latest/1.0/)
	// (https://developer.mozilla.org/en-US/docs/Web/API/WebGLRenderingContext)

	fnGetExtension             js.Value
	fnActiveTexture            js.Value
	fnAttachShader             js.Value
	fnBindBuffer               js.Value
	fnBindFramebuffer          js.Value
	fnBindTexture              js.Value
	fnBlendEquationSeparate    js.Value
	fnBlendFunc                js.Value
	fnBlendFuncSeparate        js.Value
	fnCheckFramebufferStatus   js.Value
	fnClear                    js.Value
	fnClearColor               js.Value
	fnClearDepth               js.Value
	fnClearStencil             js.Value
	fnColorMask                js.Value
	fnCompileShader            js.Value
	fnCopyTexSubImage2D        js.Value
	fnCreateBuffer             js.Value
	fnCreateFramebuffer        js.Value
	fnCreateProgram            js.Value
	fnCreateShader             js.Value
	fnCreateTexture            js.Value
	fnCullFace                 js.Value
	fnDeleteBuffer             js.Value
	fnDeleteFramebuffer        js.Value
	fnDeleteProgram            js.Value
	fnDeleteShader             js.Value
	fnDeleteTexture            js.Value
	fnDepthFunc                js.Value
	fnDepthMask                js.Value
	fnDetachShader             js.Value
	fnDisable                  js.Value
	fnDisableVertexAttribArray js.Value
	fnDrawArrays               js.Value
	fnDrawElements             js.Value
	fnEnable                   js.Value
	fnEnableVertexAttribArray  js.Value
	fnFramebufferTexture2D     js.Value
	fnGenerateMipmap           js.Value
	fnGetAttribLocation        js.Value
	fnGetError                 js.Value
	fnGetProgramParameter      js.Value
	fnGetProgramInfoLog        js.Value
	fnGetShaderParameter       js.Value
	fnGetShaderInfoLog         js.Value
	fnGetUniformLocation       js.Value
	fnLinkProgram              js.Value
	fnScissor                  js.Value
	fnShaderSource             js.Value
	fnStencilFuncSeparate      js.Value
	fnStencilOpSeparate        js.Value
	fnTexParameteri            js.Value
	fnUniform1f                js.Value
	fnUniform2f                js.Value
	fnUniform3f                js.Value
	fnUniform4f                js.Value
	fnUniform1i                js.Value
	fnUniform2i                js.Value
	fnUniform3i                js.Value
	fnUniform4i                js.Value
	fnUseProgram               js.Value
	fnVertexAttribPointer      js.Value
	fnViewport                 js.Value

	fnBufferData       js.Value
	fnBufferSubData    js.Value
	fnReadPixels       js.Value
	fnTexImage2D       js.Value
	fnTexSubImage2D    js.Value
	fnUniformMatrix4fv js.Value

	// WebGL2 functions
	// (https://www.khronos.org/registry/webgl/specs/latest/2.0/)
	// (https://developer.mozilla.org/en-US/docs/Web/API/WebGL2RenderingContext)

	fnGetBufferSubData  js.Value
	fnBlitFramebuffer   js.Value
	fnTexStorage2D      js.Value
	fnTexSubImage3D     js.Value
	fnDrawBuffers       js.Value
	fnFenceSync         js.Value
	fnDeleteSync        js.Value
	fnClientWaitSync    js.Value
	fnCreateVertexArray js.Value
	fnDeleteVertexArray js.Value
	fnBindVertexArray   js.Value
)

func initFunctions(gl js.Value) {
	fnGetExtension = getFunction(gl, "getExtension")
	fnActiveTexture = getFunction(gl, "activeTexture")
	fnAttachShader = getFunction(gl, "attachShader")
	fnBindBuffer = getFunction(gl, "bindBuffer")
	fnBindFramebuffer = getFunction(gl, "bindFramebuffer")
	fnBindTexture = getFunction(gl, "bindTexture")
	fnBlendEquationSeparate = getFunction(gl, "blendEquationSeparate")
	fnBlendFunc = getFunction(gl, "blendFunc")
	fnBlendFuncSeparate = getFunction(gl, "blendFuncSeparate")
	fnCheckFramebufferStatus = getFunction(gl, "checkFramebufferStatus")
	fnClear = getFunction(gl, "clear")
	fnClearColor = getFunction(gl, "clearColor")
	fnClearDepth = getFunction(gl, "clearDepth")
	fnClearStencil = getFunction(gl, "clearStencil")
	fnColorMask = getFunction(gl, "colorMask")
	fnCompileShader = getFunction(gl, "compileShader")
	fnCopyTexSubImage2D = getFunction(gl, "copyTexSubImage2D")
	fnCreateBuffer = getFunction(gl, "createBuffer")
	fnCreateFramebuffer = getFunction(gl, "createFramebuffer")
	fnCreateProgram = getFunction(gl, "createProgram")
	fnCreateShader = getFunction(gl, "createShader")
	fnCreateTexture = getFunction(gl, "createTexture")
	fnCullFace = getFunction(gl, "cullFace")
	fnDeleteBuffer = getFunction(gl, "deleteBuffer")
	fnDeleteFramebuffer = getFunction(gl, "deleteFramebuffer")
	fnDeleteProgram = getFunction(gl, "deleteProgram")
	fnDeleteShader = getFunction(gl, "deleteShader")
	fnDeleteTexture = getFunction(gl, "deleteTexture")
	fnDepthFunc = getFunction(gl, "depthFunc")
	fnDepthMask = getFunction(gl, "depthMask")
	fnDetachShader = getFunction(gl, "detachShader")
	fnDisable = getFunction(gl, "disable")
	fnDisableVertexAttribArray = getFunction(gl, "disableVertexAttribArray")
	fnDrawArrays = getFunction(gl, "drawArrays")
	fnDrawElements = getFunction(gl, "drawElements")
	fnEnable = getFunction(gl, "enable")
	fnEnableVertexAttribArray = getFunction(gl, "enableVertexAttribArray")
	fnFramebufferTexture2D = getFunction(gl, "framebufferTexture2D")
	fnGenerateMipmap = getFunction(gl, "generateMipmap")
	fnGetAttribLocation = getFunction(gl, "getAttribLocation")
	fnGetError = getFunction(gl, "getError")
	fnGetProgramParameter = getFunction(gl, "getProgramParameter")
	fnGetProgramInfoLog = getFunction(gl, "getProgramInfoLog")
	fnGetShaderParameter = getFunction(gl, "getShaderParameter")
	fnGetShaderInfoLog = getFunction(gl, "getShaderInfoLog")
	fnGetUniformLocation = getFunction(gl, "getUniformLocation")
	fnLinkProgram = getFunction(gl, "linkProgram")
	fnScissor = getFunction(gl, "scissor")
	fnShaderSource = getFunction(gl, "shaderSource")
	fnStencilFuncSeparate = getFunction(gl, "stencilFuncSeparate")
	fnStencilOpSeparate = getFunction(gl, "stencilOpSeparate")
	fnTexParameteri = getFunction(gl, "texParameteri")
	fnUniform1f = getFunction(gl, "uniform1f")
	fnUniform2f = getFunction(gl, "uniform2f")
	fnUniform3f = getFunction(gl, "uniform3f")
	fnUniform4f = getFunction(gl, "uniform4f")
	fnUniform1i = getFunction(gl, "uniform1i")
	fnUniform2i = getFunction(gl, "uniform2i")
	fnUniform3i = getFunction(gl, "uniform3i")
	fnUniform4i = getFunction(gl, "uniform4i")
	fnUseProgram = getFunction(gl, "useProgram")
	fnVertexAttribPointer = getFunction(gl, "vertexAttribPointer")
	fnViewport = getFunction(gl, "viewport")

	fnBufferData = getFunction(gl, "bufferData")
	fnBufferSubData = getFunction(gl, "bufferSubData")
	fnReadPixels = getFunction(gl, "readPixels")
	fnTexImage2D = getFunction(gl, "texImage2D")
	fnTexSubImage2D = getFunction(gl, "texSubImage2D")
	fnUniformMatrix4fv = getFunction(gl, "uniformMatrix4fv")

	fnGetBufferSubData = getFunction(gl, "getBufferSubData")
	fnBlitFramebuffer = getFunction(gl, "blitFramebuffer")
	fnTexStorage2D = getFunction(gl, "texStorage2D")
	fnTexSubImage3D = getFunction(gl, "texSubImage3D")
	fnDrawBuffers = getFunction(gl, "drawBuffers")
	fnFenceSync = getFunction(gl, "fenceSync")
	fnDeleteSync = getFunction(gl, "deleteSync")
	fnClientWaitSync = getFunction(gl, "clientWaitSync")
	fnCreateVertexArray = getFunction(gl, "createVertexArray")
	fnDeleteVertexArray = getFunction(gl, "deleteVertexArray")
	fnBindVertexArray = getFunction(gl, "bindVertexArray")
}

func DrawingBufferWidth() int {
	return context.Get("drawingBufferWidth").Int()
}

func DrawingBufferHeight() int {
	return context.Get("drawingBufferHeight").Int()
}

func GetExtension(name string) interface{} {
	result := fnGetExtension.Invoke(name)
	if result.IsNull() {
		return nil
	}
	// We return plain true at this point in time but in the future it might
	// be possible to return a specific extension struct here. This is why
	// the result type has been left as interface{}.
	return true
}

func ActiveTexture(texture int) {
	fnActiveTexture.Invoke(texture)
}

func AttachShader(program Program, shader Shader) {
	fnAttachShader.Invoke(js.Value(program), js.Value(shader))
}

func BindBuffer(target int, buffer Buffer) {
	fnBindBuffer.Invoke(target, js.Value(buffer))
}

func BindFramebuffer(target int, framebuffer Framebuffer) {
	fnBindFramebuffer.Invoke(target, js.Value(framebuffer))
}

func BindTexture(target int, texture Texture) {
	fnBindTexture.Invoke(target, js.Value(texture))
}

func BlendEquationSeparate(modeRGB, modeAlpha int) {
	fnBlendEquationSeparate.Invoke(modeRGB, modeAlpha)
}

func BlendFunc(sfactor, dfactor int) {
	fnBlendFunc.Invoke(sfactor, dfactor)
}

func BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha int) {
	fnBlendFuncSeparate.Invoke(srcRGB, dstRGB, srcAlpha, dstAlpha)
}

func CheckFramebufferStatus(target int) int {
	return fnCheckFramebufferStatus.Invoke(target).Int()
}

func Clear(mask int) {
	fnClear.Invoke(mask)
}

func ClearColor(r, g, b, a float32) {
	fnClearColor.Invoke(r, g, b, a)
}

func ClearDepth(depth float32) {
	fnClearDepth.Invoke(depth)
}

func ClearStencil(stencil int) {
	fnClearStencil.Invoke(stencil)
}

func ColorMask(r, g, b, a bool) {
	fnColorMask.Invoke(r, g, b, a)
}

func CompileShader(shader Shader) {
	fnCompileShader.Invoke(js.Value(shader))
}

func CopyTexSubImage2D(target, level, xoffset, yoffset, x, y, width, height int) {
	fnCopyTexSubImage2D.Invoke(target, level, xoffset, yoffset, x, y, width, height)
}

func CreateBuffer() Buffer {
	return Buffer(fnCreateBuffer.Invoke())
}

func CreateFramebuffer() Framebuffer {
	return Framebuffer(fnCreateFramebuffer.Invoke())
}

func CreateProgram() Program {
	return Program(fnCreateProgram.Invoke())
}

func CreateShader(shaderType int) Shader {
	return Shader(fnCreateShader.Invoke(shaderType))
}

func CreateTexture() Texture {
	return Texture(fnCreateTexture.Invoke())
}

func CullFace(mode int) {
	fnCullFace.Invoke(mode)
}

func DeleteBuffer(buffer Buffer) {
	fnDeleteBuffer.Invoke(js.Value(buffer))
}

func DeleteFramebuffer(framebuffer Framebuffer) {
	fnDeleteFramebuffer.Invoke(js.Value(framebuffer))
}

func DeleteProgram(program Program) {
	fnDeleteProgram.Invoke(js.Value(program))
}

func DeleteShader(shader Shader) {
	fnDeleteShader.Invoke(js.Value(shader))
}

func DeleteTexture(texture Texture) {
	fnDeleteTexture.Invoke(js.Value(texture))
}

func DepthFunc(fn int) {
	fnDepthFunc.Invoke(fn)
}

func DepthMask(mask bool) {
	fnDepthMask.Invoke(mask)
}

func DetachShader(program Program, shader Shader) {
	fnDetachShader.Invoke(js.Value(program), js.Value(shader))
}

func Disable(cap int) {
	fnDisable.Invoke(cap)
}

func DisableVertexAttribArray(index int) {
	fnDisableVertexAttribArray.Invoke(index)
}

func DrawArrays(mode, first, count int) {
	fnDrawArrays.Invoke(mode, first, count)
}

func DrawElements(mode, count, dtype, offset int) {
	fnDrawElements.Invoke(mode, count, dtype, offset)
}

func Enable(cap int) {
	fnEnable.Invoke(cap)
}

func EnableVertexAttribArray(index int) {
	fnEnableVertexAttribArray.Invoke(index)
}

func FramebufferTexture2D(target, attachment, texTarget int, texture Texture, level int) {
	fnFramebufferTexture2D.Invoke(target, attachment, texTarget, js.Value(texture), level)
}

func GenerateMipmap(target int) {
	fnGenerateMipmap.Invoke(target)
}

func GetAttribLocation(program Program, name string) AttribLocation {
	return AttribLocation(fnGetAttribLocation.Invoke(js.Value(program), name))
}

func GetError() int {
	return fnGetError.Invoke().Int()
}

func GetProgramParameter(program Program, pname int) Result {
	return Result(fnGetProgramParameter.Invoke(js.Value(program), pname))
}

func GetProgramInfoLog(program Program) string {
	return fnGetProgramInfoLog.Invoke(js.Value(program)).String()
}

func GetShaderParameter(shader Shader, pname int) Result {
	return Result(fnGetShaderParameter.Invoke(js.Value(shader), pname))
}

func GetShaderInfoLog(shader Shader) string {
	return fnGetShaderInfoLog.Invoke(js.Value(shader)).String()
}

func GetUniformLocation(program Program, name string) UniformLocation {
	return UniformLocation(fnGetUniformLocation.Invoke(js.Value(program), name))
}

func LinkProgram(program Program) {
	fnLinkProgram.Invoke(js.Value(program))
}

func Scissor(x, y, width, height int) {
	fnScissor.Invoke(x, y, width, height)
}

func ShaderSource(shader Shader, source string) {
	fnShaderSource.Invoke(js.Value(shader), source)
}

func StencilFuncSeparate(face, fun, ref, mask int) {
	fnStencilFuncSeparate.Invoke(face, fun, ref, mask)
}

func StencilOpSeparate(face, fail, zfail, zpass int) {
	fnStencilOpSeparate.Invoke(face, fail, zfail, zpass)
}

func TexParameteri(target, pname, param int) {
	fnTexParameteri.Invoke(target, pname, param)
}

func Uniform1f(location UniformLocation, x float32) {
	fnUniform1f.Invoke(js.Value(location), x)
}

func Uniform2f(location UniformLocation, x, y float32) {
	fnUniform2f.Invoke(js.Value(location), x, y)
}

func Uniform3f(location UniformLocation, x, y, z float32) {
	fnUniform3f.Invoke(js.Value(location), x, y, z)
}

func Uniform4f(location UniformLocation, x, y, z, w float32) {
	fnUniform4f.Invoke(js.Value(location), x, y, z, w)
}

func Uniform1i(location UniformLocation, x int) {
	fnUniform1i.Invoke(js.Value(location), x)
}

func Uniform2i(location UniformLocation, x, y int) {
	fnUniform2i.Invoke(js.Value(location), x, y)
}

func Uniform3i(location UniformLocation, x, y, z int) {
	fnUniform3i.Invoke(js.Value(location), x, y, z)
}

func Uniform4i(location UniformLocation, x, y, z, w int) {
	fnUniform4i.Invoke(js.Value(location), x, y, z, w)
}

func UseProgram(program Program) {
	fnUseProgram.Invoke(js.Value(program))
}

func VertexAttribPointer(index, size, dtype int, normalized bool, stride, offset int) {
	fnVertexAttribPointer.Invoke(index, size, dtype, normalized, stride, offset)
}

func Viewport(x, y, width, height int) {
	fnViewport.Invoke(x, y, width, height)
}

func BufferData(target int, data []byte, usage int) {
	// TODO: Handle nil data
	tData := newTypedSlice(data)
	defer tData.Release()
	fnBufferData.Invoke(target, tData.JSUint8Array(), usage)
}

func BufferSubData(target, dstOffset int, data []byte) {
	tData := newTypedSlice(data)
	defer tData.Release()
	fnBufferSubData.Invoke(target, dstOffset, tData.JSUint8Array())
}

func ReadPixels(x, y, width, height, format, dtype, offset int) {
	fnReadPixels.Invoke(x, y, width, height, format, dtype, offset)
}

func TexImage2D(target, level, internalFormat, width, height, border, format, dtype int, data []byte) {
	tData := newTypedSlice(data)
	defer tData.Release()
	fnTexImage2D.Invoke(target, level, internalFormat, width, height, border, format, dtype, tData.JSUint8Array())
}

func TexSubImage2D(target, level, xoffset, yoffset, width, height, format, dtype int, data []byte) {
	tData := newTypedSlice(data)
	defer tData.Release()

	switch dtype {
	case UNSIGNED_BYTE:
		buffer := tData.ArrayBuffer()
		uintArray := js.Global().Get("Uint8Array").New(buffer)
		fnTexSubImage2D.Invoke(target, level, xoffset, yoffset, width, height, format, dtype, uintArray)
	case FLOAT:
		// TODO: Optimize this with staging area
		floatArray := js.Global().Get("Float32Array").New(tData.ArrayBuffer())
		fnTexSubImage2D.Invoke(target, level, xoffset, yoffset, width, height, format, dtype, floatArray)
	default:
		panic(fmt.Errorf("unsupported dtype: %d", dtype))
	}
}

func UniformMatrix4fv(location UniformLocation, transpose bool, data []float32) {
	// TODO: Figure out a more lightweight way to do this
	jsData := make([]interface{}, len(data))
	for i, v := range data {
		jsData[i] = v
	}
	fnUniformMatrix4fv.Invoke(js.Value(location), transpose, jsData)
}

func GetBufferSubData(target, srcOffset int, data interface{}) {
	tData := newTypedSlice(data)
	defer tData.Release()
	uintArray := tData.JSUint8Array()
	fnGetBufferSubData.Invoke(target, 0, uintArray)

	sliceData := sliceToByteSlice(data)
	js.CopyBytesToGo(sliceData, uintArray)
}

func BlitFramebuffer(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter int) {
	fnBlitFramebuffer.Invoke(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter)
}

func TexStorage2D(target, levels, internalFormat, width, height int) {
	fnTexStorage2D.Invoke(target, levels, internalFormat, width, height)
}

func TexSubImage3D(target, level, xoffset, yoffset, zoffset, width, height, depth, format, dtype int, data []byte) {
	tData := newTypedSlice(data)
	defer tData.Release()
	fnTexSubImage3D.Invoke(target, level, xoffset, yoffset, zoffset, width, height, depth, format, dtype, tData.JSUint8Array())
}

func DrawBuffers(buffers []int) {
	// TODO: Figure out a more lightweight way to do this
	jsBuffers := make([]interface{}, len(buffers))
	for i, v := range buffers {
		jsBuffers[i] = v
	}
	fnDrawBuffers.Invoke(jsBuffers)
}

func FenceSync(condition, flags int) Sync {
	return Sync(fnFenceSync.Invoke(condition, flags))
}

func DeleteSync(sync Sync) {
	fnDeleteSync.Invoke(js.Value(sync))
}

func ClientWaitSync(sync Sync, flags, timeout int) int {
	return fnClientWaitSync.Invoke(js.Value(sync), flags, timeout).Int()
}

func CreateVertexArray() VertexArray {
	return VertexArray(fnCreateVertexArray.Invoke())
}

func DeleteVertexArray(array VertexArray) {
	fnDeleteVertexArray.Invoke(js.Value(array))
}

func BindVertexArray(array VertexArray) {
	fnBindVertexArray.Invoke(js.Value(array))
}
