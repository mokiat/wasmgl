//go:build js && wasm

package wasmgl

import (
	"fmt"
	"syscall/js"
	"unsafe"
)

// getFunction retrieves the function with the specified name
// from the specified target object. It returns a binding to that
// function that has target set as the function's 'this'.
func getFunction(target js.Value, name string) js.Value {
	return target.Get(name).Call("bind", target)
}

// asByteSlice returns a []byte representation for the
// specified arbitrary data type.
// It panics if the passed value cannot be represented as
// byte slice.
//
// This utility function is related to the following issues:
// https://github.com/golang/go/issues/32402
// https://github.com/golang/go/issues/31980
func asByteSlice(data interface{}) []byte {
	switch data := data.(type) {
	case []uint8:
		return data
	case []uint16:
		return unsafe.Slice((*byte)(unsafe.Pointer(&data[0])), len(data)*2)
	case []float32:
		return unsafe.Slice((*byte)(unsafe.Pointer(&data[0])), len(data)*4)
	case []float64:
		return unsafe.Slice((*byte)(unsafe.Pointer(&data[0])), len(data)*8)
	default:
		panic(fmt.Errorf("cannot convert type %T to []byte", data))
	}
}
