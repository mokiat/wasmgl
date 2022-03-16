package wasmgl

import (
	"fmt"
	"reflect"
	"syscall/js"
	"unsafe"
)

func newTypedSlice(data interface{}) *typedSlice {
	return &typedSlice{
		data: data,
	}
}

type typedSlice struct {
	data interface{} // keep reference to prevent GC
}

func (s *typedSlice) JSUint8Array() js.Value {
	byteSlice := sliceToByteSlice(s.data)
	typedArray := js.Global().Get("Uint8Array").New(len(byteSlice))
	js.CopyBytesToJS(typedArray, sliceToByteSlice(byteSlice))
	return typedArray
}

func (s *typedSlice) Release() {
	s.data = nil
}

// sliceToByteSlice is a workaround due to:
// https://github.com/golang/go/issues/32402
// https://github.com/golang/go/issues/31980
func sliceToByteSlice(data interface{}) []byte {
	switch data := data.(type) {
	case []byte:
		return data
	case []float32:
		// TODO: Consider unsafe.Slice instead
		header := (*reflect.SliceHeader)(unsafe.Pointer(&data))
		header.Len *= 4
		header.Cap *= 4
		return *(*[]byte)(unsafe.Pointer(header))
	default:
		panic(fmt.Errorf("unknown slice type %T", data))
	}
}
