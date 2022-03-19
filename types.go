//go:build js && wasm

package wasmgl

import "syscall/js"

type AttribLocation js.Value

type Buffer js.Value

type Framebuffer js.Value

var NullFramebuffer = Framebuffer(js.Null())

type Program js.Value

type Result js.Value

func (r Result) Bool() bool {
	return js.Value(r).Bool()
}

func (r Result) Int() int {
	return js.Value(r).Int()
}

type Shader js.Value

type Texture js.Value

var NullTexture = Texture(js.Null())

type UniformLocation js.Value

func (l UniformLocation) Valid() bool {
	return !js.Value(l).IsUndefined() && !js.Value(l).IsNull()
}

type VertexArray js.Value

var NullVertexArray = VertexArray(js.Null())
