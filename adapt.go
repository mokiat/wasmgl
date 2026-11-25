//go:build js && wasm

package wasmgl

import (
	"runtime"
	"slices"
	"syscall/js"
	"unsafe"
)

var (
	// NOTE: We use a global ArrayBuffer and a few TypedArray views on top
	// of it for WebGL2 calls that require such instead of allocating new ones
	// for each call.

	bufferSize   int
	arrayBuffer  js.Value
	uint8Array   js.Value
	uint16Array  js.Value
	int32Array   js.Value
	uint32Array  js.Value
	float32Array js.Value

	interfaceSlice []any
)

// DataTypes represents allowed data slice types.
type DataTypes interface {
	~int8 | ~uint8 | ~int16 | ~uint16 | ~int32 | ~uint32 | ~float32 | ~float64
}

// initialBufferSize is the size of the global ArrayBuffer when it is
// first created. It is a multiple of 8 so that all TypedArray views can
// be created over it.
const initialBufferSize = 64 * 1024

// initBuffer creates the global ArrayBuffer and its TypedArray views,
// ensuring they are valid even for calls with empty data.
func initBuffer() {
	ensureBufferSize(initialBufferSize)
}

// ensureBufferSize ensures that the global ArrayBuffer has a size
// that is equal or larger to the specified size.
//
// The buffer grows by at least a factor of two to avoid frequent
// reallocations and its size is always rounded up to a multiple of 8,
// since TypedArray views require the buffer length to be a multiple
// of their element size.
func ensureBufferSize(size int) {
	if size > bufferSize {
		newSize := max(size, 2*bufferSize)
		bufferSize = (newSize + 7) &^ 7 // round up to multiple of 8
		arrayBuffer = js.Global().Get("ArrayBuffer").New(bufferSize)
		uint8Array = js.Global().Get("Uint8Array").New(arrayBuffer)
		uint16Array = js.Global().Get("Uint16Array").New(arrayBuffer)
		int32Array = js.Global().Get("Int32Array").New(arrayBuffer)
		uint32Array = js.Global().Get("Uint32Array").New(arrayBuffer)
		float32Array = js.Global().Get("Float32Array").New(arrayBuffer)
	}
}

// pushBufferData inserts the specified data into the global
// ArrayBuffer to be used in WebGL2 calls.
//
// This function will panic if data cannot be converted to []byte
// through the asByteSlice function.
func pushBufferData[T DataTypes](data []T) {
	byteData := asByteSlice(data)
	ensureBufferSize(len(byteData))
	js.CopyBytesToJS(uint8Array, byteData)
	runtime.KeepAlive(data)
}

// popBufferData retrieves the specified data from the global
// ArrayBuffer which should have previously been used in a WebGL2
// call to be populated with data.
//
// This function will panic if data cannot be converted to []byte
// through the asByteSlice function.
func popBufferData[T DataTypes](data []T) {
	byteData := asByteSlice(data)
	js.CopyBytesToGo(byteData, uint8Array)
	runtime.KeepAlive(data)
}

// getFunction retrieves the function with the specified name
// from the specified target object. It returns a binding to that
// function that has target set as the function's 'this'.
func getFunction(target js.Value, name string) js.Value {
	return target.Get(name).Call("bind", target)
}

// asByteSlice returns a []byte representation for the
// specified arbitrary slice type.
//
// This utility function is related to the following issues:
// https://github.com/golang/go/issues/32402
// https://github.com/golang/go/issues/31980
func asByteSlice[T DataTypes](data []T) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(data))), byteSize(data))
}

// byteSize returns the number of bytes that would be
// needed to represent data once it is converted to a
// byte slice through asByteSlice.
func byteSize[T DataTypes](data []T) int {
	var zero T
	return len(data) * int(unsafe.Sizeof(zero))
}

// ensureSliceSize ensures that the global interfaceSlice has
// a size equal or larger than the specified size.
func ensureSliceSize(size int) {
	if size > len(interfaceSlice) {
		interfaceSlice = slices.Grow(interfaceSlice, size-len(interfaceSlice))[:size]
	}
}

// pushSliceData inserts the specified data into the global
// interfaceSlice at the specified offset and returns a view
// on only the inserted data.
//
// The offset parameter allows one to get two or more views over the
// global interfaceSlice if needed.
func pushSliceData[T any](data []T, offset int) []any {
	ensureSliceSize(offset + len(data))
	for i, v := range data {
		interfaceSlice[i+offset] = v
	}
	return interfaceSlice[offset : offset+len(data)]
}
