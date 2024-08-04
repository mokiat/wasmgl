//go:build js && wasm

package wasmgl

import (
	"fmt"
	"syscall/js"
)

var (
	// WebGL1 functions:
	// 	- https://www.khronos.org/registry/webgl/specs/latest/1.0/
	// 	- https://developer.mozilla.org/en-US/docs/Web/API/WebGLRenderingContext

	// WebGL2 functions
	// 	- https://www.khronos.org/registry/webgl/specs/latest/2.0/
	// 	- https://developer.mozilla.org/en-US/docs/Web/API/WebGL2RenderingContext

	// NOTE: We use references to JS functions and Invoke instead of
	// Call since the latter leads to strings being passed around
	// and TextDecoder being used on JS side.

	fnActiveTexture            js.Value
	fnAttachShader             js.Value
	fnBindBuffer               js.Value
	fnBindBufferBase           js.Value
	fnBindBufferRange          js.Value
	fnBindFramebuffer          js.Value
	fnBindSampler              js.Value
	fnBindTexture              js.Value
	fnBindVertexArray          js.Value
	fnBlendColor               js.Value
	fnBlendEquationSeparate    js.Value
	fnBlendFunc                js.Value
	fnBlendFuncSeparate        js.Value
	fnBlitFramebuffer          js.Value
	fnBufferData               js.Value
	fnBufferSubData            js.Value
	fnCheckFramebufferStatus   js.Value
	fnClear                    js.Value
	fnClearBufferfv            js.Value
	fnClearBufferiv            js.Value
	fnClearBufferuiv           js.Value
	fnClearBufferfi            js.Value
	fnClearColor               js.Value
	fnClearDepth               js.Value
	fnClearStencil             js.Value
	fnClientWaitSync           js.Value
	fnColorMask                js.Value
	fnCompileShader            js.Value
	fnCopyTexSubImage2D        js.Value
	fnCreateBuffer             js.Value
	fnCreateFramebuffer        js.Value
	fnCreateProgram            js.Value
	fnCreateSampler            js.Value
	fnCreateShader             js.Value
	fnCreateTexture            js.Value
	fnCreateVertexArray        js.Value
	fnCullFace                 js.Value
	fnDeleteBuffer             js.Value
	fnDeleteFramebuffer        js.Value
	fnDeleteProgram            js.Value
	fnDeleteSampler            js.Value
	fnDeleteShader             js.Value
	fnDeleteSync               js.Value
	fnDeleteTexture            js.Value
	fnDeleteVertexArray        js.Value
	fnDepthFunc                js.Value
	fnDepthMask                js.Value
	fnDetachShader             js.Value
	fnDisable                  js.Value
	fnDisableVertexAttribArray js.Value
	fnDrawArrays               js.Value
	fnDrawArraysInstanced      js.Value
	fnDrawBuffers              js.Value
	fnDrawElements             js.Value
	fnDrawElementsInstanced    js.Value
	fnEnable                   js.Value
	fnEnableVertexAttribArray  js.Value
	fnFinish                   js.Value
	fnFlush                    js.Value
	fnFramebufferTexture2D     js.Value
	fnFrontFace                js.Value
	fnFenceSync                js.Value
	fnGenerateMipmap           js.Value
	fnGetAttribLocation        js.Value
	fnGetBufferSubData         js.Value
	fnGetError                 js.Value
	fnGetExtension             js.Value
	fnGetParameter             js.Value
	fnGetProgramInfoLog        js.Value
	fnGetProgramParameter      js.Value
	fnGetSamplerParameter      js.Value
	fnGetShaderInfoLog         js.Value
	fnGetShaderParameter       js.Value
	fnGetSyncParameter         js.Value
	fnGetUniformBlockIndex     js.Value
	fnGetUniformLocation       js.Value
	fnInvalidateFramebuffer    js.Value
	fnIsSampler                js.Value
	fnLineWidth                js.Value
	fnLinkProgram              js.Value
	fnReadPixels               js.Value
	fnSamplerParameterf        js.Value
	fnSamplerParameteri        js.Value
	fnScissor                  js.Value
	fnShaderSource             js.Value
	fnStencilFuncSeparate      js.Value
	fnStencilMaskSeparate      js.Value
	fnStencilOpSeparate        js.Value
	fnTexImage2D               js.Value
	fnTexStorage2D             js.Value
	fnTexSubImage2D            js.Value
	fnTexSubImage3D            js.Value
	fnTexParameteri            js.Value
	fnUniform1f                js.Value
	fnUniform1i                js.Value
	fnUniform2f                js.Value
	fnUniform2i                js.Value
	fnUniform3f                js.Value
	fnUniform3i                js.Value
	fnUniform4f                js.Value
	fnUniform4i                js.Value
	fnUniformBlockBinding      js.Value
	fnUniformMatrix4fv         js.Value
	fnUseProgram               js.Value
	fnVertexAttribIPointer     js.Value
	fnVertexAttribPointer      js.Value
	fnViewport                 js.Value
)

func initFunctions(gl js.Value) {
	fnActiveTexture = getFunction(gl, "activeTexture")
	fnAttachShader = getFunction(gl, "attachShader")
	fnBindBuffer = getFunction(gl, "bindBuffer")
	fnBindBufferBase = getFunction(gl, "bindBufferBase")
	fnBindBufferRange = getFunction(gl, "bindBufferRange")
	fnBindFramebuffer = getFunction(gl, "bindFramebuffer")
	fnBindSampler = getFunction(gl, "bindSampler")
	fnBindTexture = getFunction(gl, "bindTexture")
	fnBindVertexArray = getFunction(gl, "bindVertexArray")
	fnBlendColor = getFunction(gl, "blendColor")
	fnBlendEquationSeparate = getFunction(gl, "blendEquationSeparate")
	fnBlendFunc = getFunction(gl, "blendFunc")
	fnBlendFuncSeparate = getFunction(gl, "blendFuncSeparate")
	fnBlitFramebuffer = getFunction(gl, "blitFramebuffer")
	fnBufferData = getFunction(gl, "bufferData")
	fnBufferSubData = getFunction(gl, "bufferSubData")
	fnCheckFramebufferStatus = getFunction(gl, "checkFramebufferStatus")
	fnClear = getFunction(gl, "clear")
	fnClearBufferfv = getFunction(gl, "clearBufferfv")
	fnClearBufferiv = getFunction(gl, "clearBufferiv")
	fnClearBufferuiv = getFunction(gl, "clearBufferuiv")
	fnClearBufferfi = getFunction(gl, "clearBufferfi")
	fnClearColor = getFunction(gl, "clearColor")
	fnClearDepth = getFunction(gl, "clearDepth")
	fnClearStencil = getFunction(gl, "clearStencil")
	fnClientWaitSync = getFunction(gl, "clientWaitSync")
	fnColorMask = getFunction(gl, "colorMask")
	fnCompileShader = getFunction(gl, "compileShader")
	fnCopyTexSubImage2D = getFunction(gl, "copyTexSubImage2D")
	fnCreateBuffer = getFunction(gl, "createBuffer")
	fnCreateFramebuffer = getFunction(gl, "createFramebuffer")
	fnCreateProgram = getFunction(gl, "createProgram")
	fnCreateSampler = getFunction(gl, "createSampler")
	fnCreateShader = getFunction(gl, "createShader")
	fnCreateTexture = getFunction(gl, "createTexture")
	fnCreateVertexArray = getFunction(gl, "createVertexArray")
	fnCullFace = getFunction(gl, "cullFace")
	fnDeleteBuffer = getFunction(gl, "deleteBuffer")
	fnDeleteFramebuffer = getFunction(gl, "deleteFramebuffer")
	fnDeleteProgram = getFunction(gl, "deleteProgram")
	fnDeleteSampler = getFunction(gl, "deleteSampler")
	fnDeleteShader = getFunction(gl, "deleteShader")
	fnDeleteSync = getFunction(gl, "deleteSync")
	fnDeleteTexture = getFunction(gl, "deleteTexture")
	fnDeleteVertexArray = getFunction(gl, "deleteVertexArray")
	fnDepthFunc = getFunction(gl, "depthFunc")
	fnDepthMask = getFunction(gl, "depthMask")
	fnDetachShader = getFunction(gl, "detachShader")
	fnDisable = getFunction(gl, "disable")
	fnDisableVertexAttribArray = getFunction(gl, "disableVertexAttribArray")
	fnDrawArrays = getFunction(gl, "drawArrays")
	fnDrawArraysInstanced = getFunction(gl, "drawArraysInstanced")
	fnDrawBuffers = getFunction(gl, "drawBuffers")
	fnDrawElements = getFunction(gl, "drawElements")
	fnDrawElementsInstanced = getFunction(gl, "drawElementsInstanced")
	fnEnable = getFunction(gl, "enable")
	fnEnableVertexAttribArray = getFunction(gl, "enableVertexAttribArray")
	fnFinish = getFunction(gl, "finish")
	fnFlush = getFunction(gl, "flush")
	fnFramebufferTexture2D = getFunction(gl, "framebufferTexture2D")
	fnFrontFace = getFunction(gl, "frontFace")
	fnFenceSync = getFunction(gl, "fenceSync")
	fnGenerateMipmap = getFunction(gl, "generateMipmap")
	fnGetAttribLocation = getFunction(gl, "getAttribLocation")
	fnGetBufferSubData = getFunction(gl, "getBufferSubData")
	fnGetError = getFunction(gl, "getError")
	fnGetExtension = getFunction(gl, "getExtension")
	fnGetParameter = getFunction(gl, "getParameter")
	fnGetProgramInfoLog = getFunction(gl, "getProgramInfoLog")
	fnGetProgramParameter = getFunction(gl, "getProgramParameter")
	fnGetSamplerParameter = getFunction(gl, "getSamplerParameter")
	fnGetShaderInfoLog = getFunction(gl, "getShaderInfoLog")
	fnGetShaderParameter = getFunction(gl, "getShaderParameter")
	fnGetSyncParameter = getFunction(gl, "getSyncParameter")
	fnGetUniformBlockIndex = getFunction(gl, "getUniformBlockIndex")
	fnGetUniformLocation = getFunction(gl, "getUniformLocation")
	fnInvalidateFramebuffer = getFunction(gl, "invalidateFramebuffer")
	fnIsSampler = getFunction(gl, "isSampler")
	fnLineWidth = getFunction(gl, "lineWidth")
	fnLinkProgram = getFunction(gl, "linkProgram")
	fnReadPixels = getFunction(gl, "readPixels")
	fnSamplerParameterf = getFunction(gl, "samplerParameterf")
	fnSamplerParameteri = getFunction(gl, "samplerParameteri")
	fnScissor = getFunction(gl, "scissor")
	fnShaderSource = getFunction(gl, "shaderSource")
	fnStencilFuncSeparate = getFunction(gl, "stencilFuncSeparate")
	fnStencilMaskSeparate = getFunction(gl, "stencilMaskSeparate")
	fnStencilOpSeparate = getFunction(gl, "stencilOpSeparate")
	fnTexImage2D = getFunction(gl, "texImage2D")
	fnTexStorage2D = getFunction(gl, "texStorage2D")
	fnTexSubImage2D = getFunction(gl, "texSubImage2D")
	fnTexSubImage3D = getFunction(gl, "texSubImage3D")
	fnTexParameteri = getFunction(gl, "texParameteri")
	fnUniform1f = getFunction(gl, "uniform1f")
	fnUniform1i = getFunction(gl, "uniform1i")
	fnUniform2f = getFunction(gl, "uniform2f")
	fnUniform2i = getFunction(gl, "uniform2i")
	fnUniform3f = getFunction(gl, "uniform3f")
	fnUniform3i = getFunction(gl, "uniform3i")
	fnUniform4f = getFunction(gl, "uniform4f")
	fnUniform4i = getFunction(gl, "uniform4i")
	fnUniformBlockBinding = getFunction(gl, "uniformBlockBinding")
	fnUniformMatrix4fv = getFunction(gl, "uniformMatrix4fv")
	fnUseProgram = getFunction(gl, "useProgram")
	fnVertexAttribIPointer = getFunction(gl, "vertexAttribIPointer")
	fnVertexAttribPointer = getFunction(gl, "vertexAttribPointer")
	fnViewport = getFunction(gl, "viewport")
}

func ActiveTexture(texture GLenum) {
	fnActiveTexture.Invoke(texture)
}

func AttachShader(program Program, shader Shader) {
	fnAttachShader.Invoke(js.Value(program), js.Value(shader))
}

func BindBuffer(target GLenum, buffer Buffer) {
	fnBindBuffer.Invoke(target, js.Value(buffer))
}

func BindBufferBase(target GLenum, index GLuint, buffer Buffer) {
	fnBindBufferBase.Invoke(target, index, js.Value(buffer))
}

func BindBufferRange(target GLenum, index GLuint, buffer Buffer, offset GLintptr, size GLsizeiptr) {
	fnBindBufferRange.Invoke(target, index, js.Value(buffer), offset, size)
}

func BindFramebuffer(target GLenum, framebuffer Framebuffer) {
	fnBindFramebuffer.Invoke(target, js.Value(framebuffer))
}

func BindSampler(unit GLuint, sampler Sampler) {
	fnBindSampler.Invoke(unit, js.Value(sampler))
}

func BindTexture(target GLenum, texture Texture) {
	fnBindTexture.Invoke(target, js.Value(texture))
}

func BindVertexArray(array VertexArray) {
	fnBindVertexArray.Invoke(js.Value(array))
}

func BlendColor(red, green, blue, alpha GLclampf) {
	fnBlendColor.Invoke(red, green, blue, alpha)
}

func BlendEquationSeparate(modeRGB, modeAlpha GLenum) {
	fnBlendEquationSeparate.Invoke(modeRGB, modeAlpha)
}

func BlendFunc(sfactor, dfactor GLenum) {
	fnBlendFunc.Invoke(sfactor, dfactor)
}

func BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha GLenum) {
	fnBlendFuncSeparate.Invoke(srcRGB, dstRGB, srcAlpha, dstAlpha)
}

func BlitFramebuffer(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1 GLint, mask GLbitfield, filter GLenum) {
	fnBlitFramebuffer.Invoke(srcX0, srcY0, srcX1, srcY1, dstX0, dstY0, dstX1, dstY1, mask, filter)
}

func BufferData(target GLenum, size GLsizeiptr, data []byte, usage GLenum) {
	if data != nil {
		pushBufferData(data)
		fnBufferData.Invoke(target, uint8Array, usage, 0, len(data))
	} else {
		fnBufferData.Invoke(target, size, usage)
	}
}

func BufferSubData(target GLenum, dstOffset GLintptr, data []byte) {
	pushBufferData(data)
	fnBufferSubData.Invoke(target, dstOffset, uint8Array, 0, len(data))
}

func CheckFramebufferStatus(target GLenum) GLenum {
	return GLenum(fnCheckFramebufferStatus.Invoke(target).Int())
}

func Clear(mask GLbitfield) {
	fnClear.Invoke(mask)
}

func ClearBufferfv(buffer GLenum, drawBuffer GLint, values Float32List) {
	pushBufferData(values)
	fnClearBufferfv.Invoke(buffer, drawBuffer, float32Array)
}

func ClearBufferiv(buffer GLenum, drawBuffer GLint, values Int32List) {
	pushBufferData(values)
	fnClearBufferiv.Invoke(buffer, drawBuffer, int32Array)
}

func ClearBufferuiv(buffer GLenum, drawBuffer GLint, values Uint32List) {
	pushBufferData(values)
	fnClearBufferuiv.Invoke(buffer, drawBuffer, uint32Array)
}

func ClearBufferfi(buffer GLenum, drawBuffer GLint, depth GLfloat, stencil GLint) {
	fnClearBufferfi.Invoke(buffer, drawBuffer, depth, stencil)
}

func ClearColor(r, g, b, a GLclampf) {
	fnClearColor.Invoke(r, g, b, a)
}

func ClearDepth(depth GLclampf) {
	fnClearDepth.Invoke(depth)
}

func ClearStencil(stencil GLint) {
	fnClearStencil.Invoke(stencil)
}

func ClientWaitSync(sync Sync, flags GLbitfield, timeout GLuint64) GLenum {
	return GLenum(fnClientWaitSync.Invoke(js.Value(sync), flags, timeout).Int())
}

func ColorMask(r, g, b, a GLboolean) {
	fnColorMask.Invoke(r, g, b, a)
}

func CompileShader(shader Shader) {
	fnCompileShader.Invoke(js.Value(shader))
}

func CopyTexSubImage2D(target GLenum, level, xoffset, yoffset, x, y GLint, width, height GLsizei) {
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

func CreateSampler() Sampler {
	return Sampler(fnCreateSampler.Invoke())
}

func CreateShader(shaderType GLenum) Shader {
	return Shader(fnCreateShader.Invoke(shaderType))
}

func CreateTexture() Texture {
	return Texture(fnCreateTexture.Invoke())
}

func CreateVertexArray() VertexArray {
	return VertexArray(fnCreateVertexArray.Invoke())
}

func CullFace(mode GLenum) {
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

func DeleteSampler(sampler Sampler) {
	fnDeleteSampler.Invoke(js.Value(sampler))
}

func DeleteShader(shader Shader) {
	fnDeleteShader.Invoke(js.Value(shader))
}

func DeleteSync(sync Sync) {
	fnDeleteSync.Invoke(js.Value(sync))
}

func DeleteTexture(texture Texture) {
	fnDeleteTexture.Invoke(js.Value(texture))
}

func DeleteVertexArray(array VertexArray) {
	fnDeleteVertexArray.Invoke(js.Value(array))
}

func DepthFunc(fn GLenum) {
	fnDepthFunc.Invoke(fn)
}

func DepthMask(mask GLboolean) {
	fnDepthMask.Invoke(mask)
}

func DetachShader(program Program, shader Shader) {
	fnDetachShader.Invoke(js.Value(program), js.Value(shader))
}

func Disable(cap GLenum) {
	fnDisable.Invoke(cap)
}

func DisableVertexAttribArray(index GLuint) {
	fnDisableVertexAttribArray.Invoke(index)
}

func DrawArrays(mode GLenum, first GLint, count GLsizei) {
	fnDrawArrays.Invoke(mode, first, count)
}

func DrawArraysInstanced(mode GLenum, first GLint, count, instanceCount GLsizei) {
	fnDrawArraysInstanced.Invoke(mode, first, count, instanceCount)
}

func DrawBuffers(buffers []GLenum) {
	ensureSliceSize(len(buffers))
	view := pushSliceData(buffers, 0)
	fnDrawBuffers.Invoke(view)
}

func DrawElements(mode GLenum, count GLsizei, dtype GLenum, offset GLintptr) {
	fnDrawElements.Invoke(mode, count, dtype, offset)
}

func DrawElementsInstanced(mode GLenum, count GLsizei, pType GLenum, offset GLintptr, instanceCount GLsizei) {
	fnDrawElementsInstanced.Invoke(mode, count, pType, offset, instanceCount)
}

func DrawingBufferHeight() int {
	return context.Get("drawingBufferHeight").Int()
}

func DrawingBufferWidth() int {
	return context.Get("drawingBufferWidth").Int()
}

func Enable(cap GLenum) {
	fnEnable.Invoke(cap)
}

func EnableVertexAttribArray(index GLuint) {
	fnEnableVertexAttribArray.Invoke(index)
}

func Finish() {
	fnFinish.Invoke()
}

func Flush() {
	fnFlush.Invoke()
}

func FramebufferTexture2D(target, attachment, texTarget GLenum, texture Texture, level GLint) {
	fnFramebufferTexture2D.Invoke(target, attachment, texTarget, js.Value(texture), level)
}

func FrontFace(mode GLenum) {
	fnFrontFace.Invoke(mode)
}

func FenceSync(condition GLenum, flags GLbitfield) Sync {
	return Sync(fnFenceSync.Invoke(condition, flags))
}

func GenerateMipmap(target GLenum) {
	fnGenerateMipmap.Invoke(target)
}

func GetAttribLocation(program Program, name string) GLint {
	return GLint(fnGetAttribLocation.Invoke(js.Value(program), name).Int())
}

func GetBufferSubData[T DataTypes](target GLenum, srcOffset GLintptr, data []T) {
	length := byteSize(data)
	ensureBufferSize(length)
	fnGetBufferSubData.Invoke(target, 0, uint8Array, 0, length)
	popBufferData(data)
}

func GetError() GLenum {
	return GLenum(fnGetError.Invoke().Int())
}

func GetExtension(name string) any {
	result := fnGetExtension.Invoke(name)
	if result.IsNull() {
		return nil
	}
	// We return plain true at this point in time but in the future it might
	// be possible to return a specific extension struct here. This is why
	// the result type has been left as any.
	return true
}

func GetParameter(name GLenum) Any {
	return Any(fnGetParameter.Invoke(name))
}

func GetProgramInfoLog(program Program) string {
	return fnGetProgramInfoLog.Invoke(js.Value(program)).String()
}

func GetProgramParameter(program Program, pname GLenum) Any {
	return Any(fnGetProgramParameter.Invoke(js.Value(program), pname))
}

func GetSamplerParameter(sampler Sampler, pname GLenum) Any {
	return Any(fnGetSamplerParameter.Invoke(js.Value(sampler), pname))
}

func GetShaderInfoLog(shader Shader) string {
	return fnGetShaderInfoLog.Invoke(js.Value(shader)).String()
}

func GetShaderParameter(shader Shader, pname GLenum) Any {
	return Any(fnGetShaderParameter.Invoke(js.Value(shader), pname))
}

func GetSyncParameter(sync Sync, pname GLenum) Any {
	return Any(fnGetSyncParameter.Invoke(js.Value(sync), pname))
}

func GetUniformBlockIndex(program Program, name string) GLuint {
	return GLuint(fnGetUniformBlockIndex.Invoke(js.Value(program), name).Int())
}

func GetUniformLocation(program Program, name string) UniformLocation {
	return UniformLocation(fnGetUniformLocation.Invoke(js.Value(program), name))
}

func InvalidateFramebuffer(target GLenum, attachments []GLenum) {
	ensureSliceSize(len(attachments))
	view := pushSliceData(attachments, 0)
	fnInvalidateFramebuffer.Invoke(target, view)
}

func IsSampler(sampler Sampler) bool {
	return fnIsSampler.Invoke(js.Value(sampler)).Bool()
}

func LineWidth(width GLfloat) {
	fnLineWidth.Invoke(width)
}

func LinkProgram(program Program) {
	fnLinkProgram.Invoke(js.Value(program))
}

func ReadPixels(x, y GLint, width, height GLsizei, format, dtype GLenum, offset GLintptr) {
	fnReadPixels.Invoke(x, y, width, height, format, dtype, offset)
}

func SamplerParameterf(sampler Sampler, pname GLenum, param GLfloat) {
	fnSamplerParameterf.Invoke(js.Value(sampler), pname, param)
}

func SamplerParameteri(sampler Sampler, pname GLenum, param GLint) {
	fnSamplerParameteri.Invoke(js.Value(sampler), pname, param)
}

func Scissor(x, y GLint, width, height GLsizei) {
	fnScissor.Invoke(x, y, width, height)
}

func ShaderSource(shader Shader, source string) {
	fnShaderSource.Invoke(js.Value(shader), source)
}

func StencilFuncSeparate(face, fun GLenum, ref GLint, mask GLuint) {
	fnStencilFuncSeparate.Invoke(face, fun, ref, mask)
}

func StencilMaskSeparate(face GLenum, mask GLuint) {
	fnStencilMaskSeparate.Invoke(face, mask)
}

func StencilOpSeparate(face, fail, zfail, zpass GLenum) {
	fnStencilOpSeparate.Invoke(face, fail, zfail, zpass)
}

func TexImage2D(target GLenum, level, internalFormat GLint, width, height GLsizei, border GLint, format, dtype GLenum, data []byte) {
	pushBufferData(data)
	fnTexImage2D.Invoke(target, level, internalFormat, width, height, border, format, dtype, uint8Array, 0)
}

func TexStorage2D(target GLenum, levels GLsizei, internalFormat GLenum, width, height GLsizei) {
	fnTexStorage2D.Invoke(target, levels, internalFormat, width, height)
}

func TexSubImage2D(target GLenum, level, xoffset, yoffset GLint, width, height GLsizei, format, dtype GLenum, data []byte) {
	pushBufferData(data)
	switch dtype {
	case UNSIGNED_BYTE:
		fnTexSubImage2D.Invoke(target, level, xoffset, yoffset, width, height, format, dtype, uint8Array, 0)
	case HALF_FLOAT:
		fnTexSubImage2D.Invoke(target, level, xoffset, yoffset, width, height, format, dtype, uint16Array, 0)
	case FLOAT:
		fnTexSubImage2D.Invoke(target, level, xoffset, yoffset, width, height, format, dtype, float32Array, 0)
	default:
		panic(fmt.Errorf("unsupported dtype: %d", dtype))
	}
}

func TexSubImage3D(target GLenum, level GLint, xoffset, yoffset, zoffset GLint, width, height, depth GLsizei, format, dtype GLenum, data []byte) {
	pushBufferData(data)
	fnTexSubImage3D.Invoke(target, level, xoffset, yoffset, zoffset, width, height, depth, format, dtype, uint8Array, 0)
}

func TexParameteri(target, pname GLenum, param GLint) {
	fnTexParameteri.Invoke(target, pname, param)
}

func Uniform1f(location UniformLocation, x GLfloat) {
	fnUniform1f.Invoke(js.Value(location), x)
}

func Uniform1i(location UniformLocation, x GLint) {
	fnUniform1i.Invoke(js.Value(location), x)
}

func Uniform2f(location UniformLocation, x, y GLfloat) {
	fnUniform2f.Invoke(js.Value(location), x, y)
}

func Uniform2i(location UniformLocation, x, y GLint) {
	fnUniform2i.Invoke(js.Value(location), x, y)
}

func Uniform3f(location UniformLocation, x, y, z GLfloat) {
	fnUniform3f.Invoke(js.Value(location), x, y, z)
}

func Uniform3i(location UniformLocation, x, y, z GLint) {
	fnUniform3i.Invoke(js.Value(location), x, y, z)
}

func Uniform4f(location UniformLocation, x, y, z, w GLfloat) {
	fnUniform4f.Invoke(js.Value(location), x, y, z, w)
}

func Uniform4i(location UniformLocation, x, y, z, w GLint) {
	fnUniform4i.Invoke(js.Value(location), x, y, z, w)
}

func UniformBlockBinding(program Program, index, binding GLuint) {
	fnUniformBlockBinding.Invoke(js.Value(program), index, binding)
}

func UniformMatrix4fv(location UniformLocation, transpose GLboolean, data []GLfloat) {
	pushBufferData(data)
	fnUniformMatrix4fv.Invoke(js.Value(location), transpose, float32Array, 0, len(data))
}

func UseProgram(program Program) {
	fnUseProgram.Invoke(js.Value(program))
}

func VertexAttribIPointer(index GLuint, size GLint, dtype GLenum, stride GLsizei, offset GLintptr) {
	fnVertexAttribIPointer.Invoke(index, size, dtype, stride, offset)
}

func VertexAttribPointer(index GLuint, size GLint, dtype GLenum, normalized GLboolean, stride GLsizei, offset GLintptr) {
	fnVertexAttribPointer.Invoke(index, size, dtype, normalized, stride, offset)
}

func Viewport(x, y GLint, width, height GLsizei) {
	fnViewport.Invoke(x, y, width, height)
}
