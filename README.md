# Go WASM GL

This project aims to expose the WebGL2 API as Go API that can be used
in wasm projects.

> **Warning:** The project is in early development and the API is likely to
> change in backward incompatible ways!

## Getting Started

You need to add the project as a dependency.

```
go get github.com/mokiat/wasmgl@latest
```

The implementation uses `syscall/js` calls and as such requires that client
applications are compiled with the `GOOS=js` and `GOARCH=wasm` options.

If you are unfamiliar with how Go and WASM works, then you should have a look at
the official [WebAssembly with Go documentation](https://github.com/golang/go/wiki/WebAssembly).

Following is an example Go code that clears the canvas with a green color.

```go
package main

import (
	"log"

	"github.com/mokiat/wasmgl"
)

func main() {
	const canvasElementID = "glcanvas"

	if err := wasmgl.InitFromID(canvasElementID); err != nil {
		log.Fatalf("Failed to initialize wasmgl: %v", err)
	}

	wasmgl.ClearColor(0.0, 1.0, 0.0, 1.0)
	wasmgl.Clear(wasmgl.COLOR_BUFFER_BIT)
}
```
