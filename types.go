//go:build js && wasm

package wasmgl

import "syscall/js"

type Shader js.Value

type Texture js.Value

type VertexArray js.Value
