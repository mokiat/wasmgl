//go:build js && wasm

package wasmgl

import "syscall/js"

var (
	// ARRAY_BUFFER                int
	// CLAMP_TO_EDGE               int
	COLOR_BUFFER_BIT int
	// COMPILE_STATUS              int
	// CULL_FACE                   int
	// DEPTH_BUFFER_BIT            int
	// DEPTH_TEST                  int
	// ELEMENT_ARRAY_BUFFER        int
	// FLOAT                       int
	// FRAGMENT_SHADER             int
	// LEQUAL                      int
	// LESS                        int
	// LINEAR                      int
	// LINK_STATUS                 int
	// RGB                         int
	// RGBA                        int
	// STATIC_DRAW                 int
	// TEXTURE0                    int
	// TEXTURE1                    int
	// TEXTURE2                    int
	// TEXTURE_2D                  int
	// TEXTURE_CUBE_MAP            int
	// TEXTURE_CUBE_MAP_POSITIVE_X int
	// TEXTURE_CUBE_MAP_NEGATIVE_X int
	// TEXTURE_CUBE_MAP_POSITIVE_Y int
	// TEXTURE_CUBE_MAP_NEGATIVE_Y int
	// TEXTURE_CUBE_MAP_POSITIVE_Z int
	// TEXTURE_CUBE_MAP_NEGATIVE_Z int
	// TEXTURE_MAG_FILTER          int
	// TEXTURE_MIN_FILTER          int
	// TEXTURE_WRAP_S              int
	// TEXTURE_WRAP_T              int
	// TRIANGLES                   int
	// UNSIGNED_BYTE               int
	// UNSIGNED_SHORT              int
	// VERTEX_SHADER               int
)

func initConstants(gl js.Value) {
	// ARRAY_BUFFER:                gl.Get("ARRAY_BUFFER").Int()
	// CLAMP_TO_EDGE:               gl.Get("CLAMP_TO_EDGE").Int()
	COLOR_BUFFER_BIT = gl.Get("COLOR_BUFFER_BIT").Int()
	// COMPILE_STATUS:              gl.Get("COMPILE_STATUS").Int()
	// CULL_FACE:                   gl.Get("CULL_FACE").Int()
	// DEPTH_BUFFER_BIT:            gl.Get("DEPTH_BUFFER_BIT").Int()
	// DEPTH_TEST:                  gl.Get("DEPTH_TEST").Int()
	// ELEMENT_ARRAY_BUFFER:        gl.Get("ELEMENT_ARRAY_BUFFER").Int()
	// FLOAT:                       gl.Get("FLOAT").Int()
	// FRAGMENT_SHADER:             gl.Get("FRAGMENT_SHADER").Int()
	// LEQUAL:                      gl.Get("LEQUAL").Int()
	// LESS:                        gl.Get("LESS").Int()
	// LINEAR:                      gl.Get("LINEAR").Int()
	// LINK_STATUS:                 gl.Get("LINK_STATUS").Int()
	// RGB:                         gl.Get("RGB").Int()
	// RGBA:                        gl.Get("RGBA").Int()
	// STATIC_DRAW:                 gl.Get("STATIC_DRAW").Int()
	// TEXTURE0:                    gl.Get("TEXTURE0").Int()
	// TEXTURE1:                    gl.Get("TEXTURE1").Int()
	// TEXTURE2:                    gl.Get("TEXTURE2").Int()
	// TEXTURE_2D:                  gl.Get("TEXTURE_2D").Int()
	// TEXTURE_CUBE_MAP:            gl.Get("TEXTURE_CUBE_MAP").Int()
	// TEXTURE_CUBE_MAP_POSITIVE_X: gl.Get("TEXTURE_CUBE_MAP_POSITIVE_X").Int()
	// TEXTURE_CUBE_MAP_NEGATIVE_X: gl.Get("TEXTURE_CUBE_MAP_NEGATIVE_X").Int()
	// TEXTURE_CUBE_MAP_POSITIVE_Y: gl.Get("TEXTURE_CUBE_MAP_POSITIVE_Y").Int()
	// TEXTURE_CUBE_MAP_NEGATIVE_Y: gl.Get("TEXTURE_CUBE_MAP_NEGATIVE_Y").Int()
	// TEXTURE_CUBE_MAP_POSITIVE_Z: gl.Get("TEXTURE_CUBE_MAP_POSITIVE_Z").Int()
	// TEXTURE_CUBE_MAP_NEGATIVE_Z: gl.Get("TEXTURE_CUBE_MAP_NEGATIVE_Z").Int()
	// TEXTURE_MAG_FILTER:          gl.Get("TEXTURE_MAG_FILTER").Int()
	// TEXTURE_MIN_FILTER:          gl.Get("TEXTURE_MIN_FILTER").Int()
	// TEXTURE_WRAP_S:              gl.Get("TEXTURE_WRAP_S").Int()
	// TEXTURE_WRAP_T:              gl.Get("TEXTURE_WRAP_T").Int()
	// TRIANGLES:                   gl.Get("TRIANGLES").Int()
	// UNSIGNED_BYTE:               gl.Get("UNSIGNED_BYTE").Int()
	// UNSIGNED_SHORT:              gl.Get("UNSIGNED_SHORT").Int()
	// VERTEX_SHADER:               gl.Get("VERTEX_SHADER").Int()
}
