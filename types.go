//go:build js && wasm

package wasmgl

import "syscall/js"

var NilAttribLocation = AttribLocation(js.Null())

type AttribLocation js.Value

func (l AttribLocation) IsValid() bool {
	return isSpecified(js.Value(l))
}

var NilBuffer = Buffer(js.Null())

type Buffer js.Value

func (b Buffer) IsValid() bool {
	return isSpecified(js.Value(b))
}

var NilFramebuffer = Framebuffer(js.Null())

type Framebuffer js.Value

func (f Framebuffer) IsValid() bool {
	return isSpecified(js.Value(f))
}

var NilProgram = Program(js.Null())

type Program js.Value

func (p Program) IsValid() bool {
	return isSpecified(js.Value(p))
}

var NilResult = Result(js.Null())

type Result js.Value

func (r Result) IsValid() bool {
	return isSpecified(js.Value(r))
}

func (r Result) Bool() bool {
	return js.Value(r).Bool()
}

func (r Result) Int() int {
	return js.Value(r).Int()
}

var NilShader = Shader(js.Null())

type Shader js.Value

func (s Shader) IsValid() bool {
	return isSpecified(js.Value(s))
}

var NilTexture = Texture(js.Null())

type Texture js.Value

func (t Texture) IsValid() bool {
	return isSpecified(js.Value(t))
}

var NilUniformLocation = UniformLocation(js.Null())

type UniformLocation js.Value

func (l UniformLocation) IsValid() bool {
	return isSpecified(js.Value(l))
}

var NilVertexArray = VertexArray(js.Null())

type VertexArray js.Value

func (a VertexArray) IsValid() bool {
	return isSpecified(js.Value(a))
}

var NilSync = Sync(js.Null())

type Sync js.Value

func (s Sync) Valid() bool {
	return isSpecified(js.Value(s))
}

func isSpecified(jsValue js.Value) bool {
	return !jsValue.IsUndefined() && !jsValue.IsNull()
}
