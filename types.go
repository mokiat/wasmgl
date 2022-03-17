//go:build js && wasm

package wasmgl

import "syscall/js"

type AttribLocation js.Value

type Buffer js.Value

type Framebuffer js.Value

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

type UniformLocation js.Value

func (l UniformLocation) Valid() bool {
	return !js.Value(l).IsUndefined()
}

type VertexArray js.Value

var NullVertexArray = VertexArray(js.Null())
