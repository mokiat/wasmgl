//go:build js && wasm

package wasmgl

import (
	"syscall/js"
)

// TODO: Use one big ArrayBuffer for all data transfers instead of allocating
// each time.

func newTypedSlice(data interface{}) *typedSlice {
	return &typedSlice{
		data: data,
	}
}

type typedSlice struct {
	data interface{} // keep reference to prevent GC
}

func (s *typedSlice) JSUint8Array() js.Value {
	byteSlice := asByteSlice(s.data)
	typedArray := js.Global().Get("Uint8Array").New(len(byteSlice))
	js.CopyBytesToJS(typedArray, byteSlice)
	return typedArray
}

func (s *typedSlice) ArrayBuffer() js.Value {
	byteSlice := asByteSlice(s.data)
	buffer := js.Global().Get("ArrayBuffer").New(len(byteSlice))

	typedArray := js.Global().Get("Uint8Array").New(buffer)
	js.CopyBytesToJS(typedArray, byteSlice)

	return buffer
}

func (s *typedSlice) Release() {
	s.data = nil
}
