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
	// GLint64 represents the GLint64 type from the specification.
	GLint64 = int64
	// GLuint64 represents the GLuint64 type from the specification.
	GLuint64 = uint64

	// Float32List represents the Float32List type from the specification.
	Float32List = []float32
	// Int32List represents the Int32List type from the specification.
	Int32List = []int32
	// Uint32List represents the Uint32List type from the specification.
	Uint32List = []uint32
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

// Result is a legacy alias for Any.
//
// Deprecated: Use Any instead.
type Result = Any

// Any represents an undefined return type. In the specification this is
// indicated with the `any` keyword.
type Any js.Value

// IsValid returns whether this Any is specified and can be used.
func (r Any) IsValid() bool {
	return isSpecified(js.Value(r))
}

// GLboolean returns the contents of this Any as a GLboolean type.
func (r Any) GLboolean() GLboolean {
	return js.Value(r).Bool()
}

// GLenum returns the contents of this Any as a GLenum type.
func (r Any) GLenum() GLenum {
	return GLenum(js.Value(r).Int())
}

// GLint returns the contents of this Any as a GLint type.
func (r Any) GLint() GLint {
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

// NilSampler equals the zero Sampler.
var NilSampler = Sampler(js.Null())

// Sampler represents the WebGLSampler type from the specification.
type Sampler js.Value

// IsValid returns whether this Sampler is different from the zero Sampler or
// an unspecified Sampler.
func (s Sampler) IsValid() bool {
	return isSpecified(js.Value(s))
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
