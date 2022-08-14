//go:build js && wasm

package wasmgl

import "syscall/js"

type (
	// GLenum represents the GLenum type from the specification.
	GLenum = uint32
	// GLboolean represents the GLboolean type from the specification.
	GLboolean = bool
	// GLbitfield represents the GLbitfield type from the specification.
	GLbitfield = uint32
	// GLbyte represents the GLbyte type from the specification.
	GLbyte = int8
	// GLshort represents the GLshort type from the specification.
	GLshort = uint16
	// GLint represents the GLint type from the specification.
	GLint = int32
	// GLsizei represents the GLsizei type from the specification.
	GLsizei = int32
	// GLintptr represents the GLintptr type from the specification.
	GLintptr = int64
	// GLsizeiptr represents the GLsizeiptr type from the specification.
	GLsizeiptr = int64
	// GLubyte represents the GLubyte type from the specification.
	GLubyte = uint8
	// GLushort represents the GLushort type from the specification.
	GLushort = uint16
	// GLuint represents the GLuint type from the specification.
	GLuint = uint32
	// GLfloat represents the GLfloat type from the specification.
	GLfloat = float32
	// GLclampf represents the GLclampf type from the specification.
	GLclampf = float32
)

// NilBuffer equals the zero Buffer.
var NilBuffer = Buffer(js.Null())

// Buffer represents the WebGLBuffer type from the specification.
type Buffer js.Value

// IsValid returns whether this Buffer is different from the zero Buffer
// or an unspecified Buffer.
func (b Buffer) IsValid() bool {
	return isSpecified(js.Value(b))
}

// NilFramebuffer equals the zero Framebuffer.
var NilFramebuffer = Framebuffer(js.Null())

// Framebuffer represents the WebGLFramebuffer type from the specification.
type Framebuffer js.Value

// IsValid returns whether this Framebuffer is different from the zero
// Framebuffer or an unspecified Framebuffer.
func (f Framebuffer) IsValid() bool {
	return isSpecified(js.Value(f))
}

// NilProgram equals the zero Program.
var NilProgram = Program(js.Null())

// Program represents the WebGLProgram type from the specification.
type Program js.Value

// IsValid returns whether this Program is different from the zero Program or
// an unspecified Program.
func (p Program) IsValid() bool {
	return isSpecified(js.Value(p))
}

// Result represents an undefined return type. In the specification this is
// indicated with the `any` keyword.
type Result js.Value

// IsValid returns whether this Result is specified and can be used.
func (r Result) IsValid() bool {
	return isSpecified(js.Value(r))
}

// GLboolean returns the contents of this Result as a GLboolean type.
func (r Result) GLboolean() GLboolean {
	return js.Value(r).Bool()
}

// GLenum returns the contents of this Result as a GLenum type.
func (r Result) GLenum() GLenum {
	return GLenum(js.Value(r).Int())
}

// GLint returns the contents of this Result as a GLint type.
func (r Result) GLint() GLint {
	return GLint(js.Value(r).Int())
}

// NilShader equals the zero Shader.
var NilShader = Shader(js.Null())

// Shader represents the WebGLShader type from the specification.
type Shader js.Value

// IsValid returns whether this Shader is different from the zero Shader or
// an unspecified Shader.
func (s Shader) IsValid() bool {
	return isSpecified(js.Value(s))
}

// NilSync equals the zero Sync.
var NilSync = Sync(js.Null())

// Sync represents the WebGLSync type from the specification.
type Sync js.Value

// IsValid returns whether this Sync is different from the zero Sync or an
// unspecified Sync.
func (s Sync) Valid() bool {
	return isSpecified(js.Value(s))
}

// NilTexture equals the zero Texture.
var NilTexture = Texture(js.Null())

// Texture represents the WebGLTexture type from the specification.
type Texture js.Value

// IsValid returns whether this Texture is different from the zero Texture or
// an unspecified Texture.
func (t Texture) IsValid() bool {
	return isSpecified(js.Value(t))
}

// NilUniformLocation equals the nil UniformLocation.
var NilUniformLocation = UniformLocation(js.Null())

// UniformLocation represents the WebGLUniformLocation type from the
// specification.
type UniformLocation js.Value

// IsValid returns whether this UniformLocation is different from the nil
// UniformLocation or an unspecified UniformLocation.
func (l UniformLocation) IsValid() bool {
	return isSpecified(js.Value(l))
}

// NilVertexArray equals the zero VertexArray.
var NilVertexArray = VertexArray(js.Null())

// VertexArray represents the WebGLVertexArrayObject type from the
// specification.
type VertexArray js.Value

// IsValid returns whether this VertexArray is different from the zero
// VertexArray or an unspecified VertexArray.
func (a VertexArray) IsValid() bool {
	return isSpecified(js.Value(a))
}

func isSpecified(jsValue js.Value) bool {
	return !jsValue.IsUndefined() && !jsValue.IsNull()
}
