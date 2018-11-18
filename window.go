// +build js,wasm

package wasmgl

import (
	"fmt"
	"syscall/js"
	"time"
)

// WindowHints structure is used to specify creation hints to the CreateWindow
// function.
type WindowHints struct {
	// Width specifies the desired width of the canvas element. This is applied
	// only if Fullscreen is false and Height is set to non-zero value
	Width int

	// Height specifies the desired height of the canvas element. This is applied
	// only if Fullscreen is false and Width is set to non-zero value
	Height int

	// Fullscreen specifies whether the canvas should cover the whole body of the
	// HTML page.
	Fullscreen bool
}

// CreateWindow creates a Window instance that works on an HTML canvas element and
// abstracts away some OpenGL bootstrap calls.
//
// Once the Window instance is no longer needed, the Destroy method should be called
// to release any allocated resources.
func CreateWindow(canvasID string, hints WindowHints) (*Window, error) {
	htmlDocument := js.Global().Get("document")
	if htmlDocument == js.Undefined() {
		return nil, fmt.Errorf("could not locate document element")
	}

	htmlCanvas := htmlDocument.Call("getElementById", canvasID)
	if htmlCanvas == js.Undefined() {
		return nil, fmt.Errorf("could not locate canvas element")
	}

	if hints.Fullscreen {
		htmlBody := htmlDocument.Get("body")
		if htmlBody == js.Undefined() {
			return nil, fmt.Errorf("could not locate body element")
		}
		bodyWidth := htmlBody.Get("clientWidth").Int()
		bodyHeight := htmlBody.Get("clientHeight").Int()
		htmlCanvas.Set("width", bodyWidth)
		htmlCanvas.Set("height", bodyHeight)
	} else if hints.Width != 0 && hints.Height != 0 {
		htmlCanvas.Set("width", hints.Width)
		htmlCanvas.Set("height", hints.Height)
	}

	return &Window{
		htmlCanvas: htmlCanvas,
	}, nil
}

// Window abstracts an HTML canvas element and any
// bootstrap logic.
type Window struct {
	htmlCanvas js.Value
}

// InitGL creates a new GL (WebGL1) instance for the given Window or returns
// an error on failure.
func (w *Window) InitGL() (*GL, error) {
	return createGL(w.htmlCanvas)
}

// InitGL2 creates a new GL2 (WebGL2) instance for the given Window or returns
// an error on failure.
func (w *Window) InitGL2() (*GL2, error) {
	return createGL2(w.htmlCanvas)
}

type LoopCallback func(elapsedSeconds float32)

// Loop initiates an update loop. The specified callback will be called
// to perform any scene rendering.
func (w *Window) Loop(delegate LoopCallback) {
	var callback js.Callback
	lastTick := time.Now()
	// FIXME: callback is never released
	callback = js.NewCallback(func(args []js.Value) {
		elapsedSeconds := time.Now().Sub(lastTick).Seconds()
		delegate(float32(elapsedSeconds))
		lastTick = time.Now()

		js.Global().Call("requestAnimationFrame", callback)
	})
	js.Global().Call("requestAnimationFrame", callback)
}

// Destroy releases any allocated resources.
func (w *Window) Destroy() {
}
