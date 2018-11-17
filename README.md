# WASM GL

This project aims at simplifying WebGL usage from Go WebAssembly modules.

Normally a developer would be required to use `syscall/js` to manually call into JS objects in order to issue WebGL calls. The functions and types exposed by this project abstract away this and allow for native GL invocations (which instead do all the `syscall/js` calls).

> **Warning: The project is in early development by me, during my free time, and the API is likely to change in backward incompatible way.**

## Getting Started

You need to add the project as a dependency.

```
go get github.com/mokiat/wasmgl
```

When developing in your editor (e.g. Visual Studio Code), you need to make sure you have started it with the proper environment variables set, otherwise you will get error messages. This is related to the `syscall/js` package.

```bash
export GOOS=js
export GOARCH=wasm
```

You should follow the official [Go WebAssembly documentation](https://github.com/golang/go/wiki/WebAssembly) on how to set up an HTML project.

Following is an example on how your HTML could look like.

```html
<html>
    <head>
        <meta charset="utf-8">
        <script src="wasm_exec.js"></script>
        <script>
            const go = new Go();
            WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
                go.run(result.instance);
            });
        </script>
    </head>
    <body>
        <canvas id="glcanvas"></canvas>
    </body>
</html>
```

Following is an example Go code that fills the canvas with green color though WebGL.

```go
package main

import (
	"log"

	"github.com/mokiat/wasmgl"
)

const (
	width  = 800
	height = 600
)

func main() {
	window, err := wasmgl.CreateWindow("glcanvas", wasmgl.WindowHints{
		Width:  width,
		Height: height,
	})
	if err != nil {
		log.Fatalf("failed to initialize window: %s", err)
	}

	gl, err := window.InitGL2()
	if err != nil {
		log.Fatalf("failed to initialize webgl2: %s", err)
	}
	gl.Viewport(0, 0, width, height)

	window.Loop(func(elaspedSeconds float32) {
		gl.ClearColor(0.0, 1.0, 0.0, 1.0)
		gl.Clear(gl.ColorBufferBit)
	})

	<-make(chan struct{})
}
```
