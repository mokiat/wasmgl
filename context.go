//go:build js && wasm

package wasmgl

import (
	"fmt"
	"syscall/js"
)

var context js.Value

// ContextOption represents a configuration option for the WebGL2 context.
type ContextOption func(v js.Value)

// WithOptionAlpha configures the alpha context option.
func WithOptionAlpha(alpha bool) ContextOption {
	return func(v js.Value) {
		v.Set("alpha", alpha)
	}
}

// WithOptionDepth configures the depth context option.
func WithOptionDepth(depth bool) ContextOption {
	return func(v js.Value) {
		v.Set("depth", depth)
	}
}

// WithOptionStencil configures the stencil context option.
func WithOptionStencil(stencil bool) ContextOption {
	return func(v js.Value) {
		v.Set("stencil", stencil)
	}
}

// WithOptionDesynchronized configures the alpha context option.
func WithOptionDesynchronized(desynchronized bool) ContextOption {
	return func(v js.Value) {
		v.Set("desynchronized", desynchronized)
	}
}

// WithOptionAntialias configures the antialias context option.
func WithOptionAntialias(antialias bool) ContextOption {
	return func(v js.Value) {
		v.Set("antialias", antialias)
	}
}

// WithOptionFailIfMajorPerformanceCaveat configures the failIfMajorPerformanceCaveat context option.
func WithOptionFailIfMajorPerformanceCaveat(fail bool) ContextOption {
	return func(v js.Value) {
		v.Set("failIfMajorPerformanceCaveat", fail)
	}
}

// PowerPreference represents a context power preference.
type PowerPreference string

const (
	PowerPreferenceDefault         PowerPreference = "default"
	PowerPreferenceHighPerformance PowerPreference = "high-performance"
	PowerPreferenceLowPower        PowerPreference = "low-power"
)

// WithOptionPowerPreference configures the powerPreference context option.
func WithOptionPowerPreference(pref PowerPreference) ContextOption {
	return func(v js.Value) {
		v.Set("powerPreference", string(pref))
	}
}

// WithOptionPremultipliedAlpha configures the premultipliedAlpha context option.
func WithOptionPremultipliedAlpha(premultiplied bool) ContextOption {
	return func(v js.Value) {
		v.Set("premultipliedAlpha", premultiplied)
	}
}

// WithOptionPreserveDrawingBuffer configures the preserveDrawingBuffer context option.
func WithOptionPreserveDrawingBuffer(preserve bool) ContextOption {
	return func(v js.Value) {
		v.Set("preserveDrawingBuffer", preserve)
	}
}

// WithOptionXRCompatible configures the xrCompatible context option.
func WithOptionXRCompatible(compatible bool) ContextOption {
	return func(v js.Value) {
		v.Set("xrCompatible", compatible)
	}
}

// InitFromID initializes webgl context and bindings
// from the canvas that has the specified canvasID ID.
func InitFromID(canvasID string, opts ...ContextOption) error {
	htmlDocument := js.Global().Get("document")
	if htmlDocument.IsUndefined() {
		return fmt.Errorf("could not locate document element")
	}
	htmlCanvas := htmlDocument.Call("getElementById", canvasID)
	if htmlCanvas.IsNull() {
		return fmt.Errorf("could not locate canvas element with id %q", canvasID)
	}
	return InitFromCanvas(htmlCanvas, opts...)
}

// InitFromCanvas initializes webgl context and bindings
// from the specified htmlCanvas canvas element reference.
func InitFromCanvas(htmlCanvas js.Value, opts ...ContextOption) error {
	optsObject := js.Global().Get("Object").New()
	for _, opt := range opts {
		opt(optsObject)
	}
	context = htmlCanvas.Call("getContext", "webgl2", optsObject)
	if context.IsNull() {
		return fmt.Errorf("could not acquire webgl2 context")
	}
	initFunctions(context)
	return nil
}
