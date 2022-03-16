//go:build js && wasm

package wasmgl

import "syscall/js"

type Buffer js.Value

type Program js.Value

type Shader js.Value

type Texture js.Value

type VertexArray js.Value
