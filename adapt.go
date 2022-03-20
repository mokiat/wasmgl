//go:build js && wasm

package wasmgl

import "syscall/js"

// getFunction retrieves the function with the specified name
// from the specified target object. It returns a binding to that
// function that has target set as the function's 'this'.
func getFunction(target js.Value, name string) js.Value {
	return target.Get(name).Call("bind", target)
}
