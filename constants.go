//go:build js && wasm

package wasmgl

import "syscall/js"

// Reference on WebGL2 constants:
// https://developer.mozilla.org/en-US/docs/Web/API/WebGL_API/Constants

var (
	BACK                        int
	BLEND                       int
	CLAMP_TO_EDGE               int
	COLOR_BUFFER_BIT            int
	COMPILE_STATUS              int
	CULL_FACE                   int
	DEPTH_BUFFER_BIT            int
	DEPTH_TEST                  int
	FLOAT                       int
	FRAGMENT_SHADER             int
	FRONT                       int
	FRONT_AND_BACK              int
	LINEAR                      int
	LINEAR_MIPMAP_LINEAR        int
	LINEAR_MIPMAP_NEAREST       int
	LINK_STATUS                 int
	NEAREST                     int
	NEAREST_MIPMAP_LINEAR       int
	NEAREST_MIPMAP_NEAREST      int
	ONE_MINUS_SRC_ALPHA         int
	REPEAT                      int
	SRC_ALPHA                   int
	SRGB8_ALPHA8                int
	STENCIL_BUFFER_BIT          int
	STENCIL_TEST                int
	TEXTURE0                    int
	TEXTURE1                    int
	TEXTURE2                    int
	TEXTURE_2D                  int
	TEXTURE_CUBE_MAP            int
	TEXTURE_CUBE_MAP_POSITIVE_X int
	TEXTURE_CUBE_MAP_NEGATIVE_X int
	TEXTURE_CUBE_MAP_POSITIVE_Y int
	TEXTURE_CUBE_MAP_NEGATIVE_Y int
	TEXTURE_CUBE_MAP_POSITIVE_Z int
	TEXTURE_CUBE_MAP_NEGATIVE_Z int
	TEXTURE_MAG_FILTER          int
	TEXTURE_MIN_FILTER          int
	TEXTURE_WRAP_R              int
	TEXTURE_WRAP_S              int
	TEXTURE_WRAP_T              int
	TRIANGLES                   int
	TRIANGLE_FAN                int
	UNSIGNED_BYTE               int
	UNSIGNED_SHORT              int
	VERTEX_SHADER               int

	// buffers
	STATIC_DRAW          int
	STREAM_DRAW          int
	DYNAMIC_DRAW         int
	ARRAY_BUFFER         int
	ELEMENT_ARRAY_BUFFER int
	BUFFER_SIZE          int
	BUFFER_USAGE         int

	// depth or stencil tests
	NEVER    int
	LESS     int
	EQUAL    int
	LEQUAL   int
	GREATER  int
	NOTEQUAL int
	GEQUAL   int
	ALWAYS   int

	// framebuffers and renderbuffers
	FRAMEBUFFER    int
	STENCIL_INDEX8 int

	// stencil operations
	KEEP      int
	REPLACE   int
	INCR      int
	DECR      int
	INVERT    int
	INCR_WRAP int
	DECR_WRAP int

	// textures
	R8    int
	RGB   int
	RGBA  int
	RGBA8 int
)

func initConstants(gl js.Value) {
	BLEND = gl.Get("BLEND").Int()
	BACK = gl.Get("BACK").Int()
	CLAMP_TO_EDGE = gl.Get("CLAMP_TO_EDGE").Int()
	COLOR_BUFFER_BIT = gl.Get("COLOR_BUFFER_BIT").Int()
	COMPILE_STATUS = gl.Get("COMPILE_STATUS").Int()
	CULL_FACE = gl.Get("CULL_FACE").Int()
	DEPTH_BUFFER_BIT = gl.Get("DEPTH_BUFFER_BIT").Int()
	DEPTH_TEST = gl.Get("DEPTH_TEST").Int()
	FLOAT = gl.Get("FLOAT").Int()
	FRAGMENT_SHADER = gl.Get("FRAGMENT_SHADER").Int()
	FRONT = gl.Get("FRONT").Int()
	FRONT_AND_BACK = gl.Get("FRONT_AND_BACK").Int()
	LINEAR = gl.Get("LINEAR").Int()
	LINEAR_MIPMAP_LINEAR = gl.Get("LINEAR_MIPMAP_LINEAR").Int()
	LINEAR_MIPMAP_NEAREST = gl.Get("LINEAR_MIPMAP_NEAREST").Int()
	LINK_STATUS = gl.Get("LINK_STATUS").Int()
	NEAREST = gl.Get("NEAREST").Int()
	NEAREST_MIPMAP_LINEAR = gl.Get("NEAREST_MIPMAP_LINEAR").Int()
	NEAREST_MIPMAP_NEAREST = gl.Get("NEAREST_MIPMAP_NEAREST").Int()
	ONE_MINUS_SRC_ALPHA = gl.Get("ONE_MINUS_SRC_ALPHA").Int()
	REPEAT = gl.Get("REPEAT").Int()
	SRC_ALPHA = gl.Get("SRC_ALPHA").Int()
	SRGB8_ALPHA8 = gl.Get("SRGB8_ALPHA8").Int()
	STENCIL_BUFFER_BIT = gl.Get("STENCIL_BUFFER_BIT").Int()
	STENCIL_TEST = gl.Get("STENCIL_TEST").Int()
	TEXTURE0 = gl.Get("TEXTURE0").Int()
	TEXTURE1 = gl.Get("TEXTURE1").Int()
	TEXTURE2 = gl.Get("TEXTURE2").Int()
	TEXTURE_2D = gl.Get("TEXTURE_2D").Int()
	TEXTURE_CUBE_MAP = gl.Get("TEXTURE_CUBE_MAP").Int()
	TEXTURE_CUBE_MAP_POSITIVE_X = gl.Get("TEXTURE_CUBE_MAP_POSITIVE_X").Int()
	TEXTURE_CUBE_MAP_NEGATIVE_X = gl.Get("TEXTURE_CUBE_MAP_NEGATIVE_X").Int()
	TEXTURE_CUBE_MAP_POSITIVE_Y = gl.Get("TEXTURE_CUBE_MAP_POSITIVE_Y").Int()
	TEXTURE_CUBE_MAP_NEGATIVE_Y = gl.Get("TEXTURE_CUBE_MAP_NEGATIVE_Y").Int()
	TEXTURE_CUBE_MAP_POSITIVE_Z = gl.Get("TEXTURE_CUBE_MAP_POSITIVE_Z").Int()
	TEXTURE_CUBE_MAP_NEGATIVE_Z = gl.Get("TEXTURE_CUBE_MAP_NEGATIVE_Z").Int()
	TEXTURE_MAG_FILTER = gl.Get("TEXTURE_MAG_FILTER").Int()
	TEXTURE_MIN_FILTER = gl.Get("TEXTURE_MIN_FILTER").Int()
	TEXTURE_WRAP_R = gl.Get("TEXTURE_WRAP_R").Int()
	TEXTURE_WRAP_S = gl.Get("TEXTURE_WRAP_S").Int()
	TEXTURE_WRAP_T = gl.Get("TEXTURE_WRAP_T").Int()
	TRIANGLES = gl.Get("TRIANGLES").Int()
	TRIANGLE_FAN = gl.Get("TRIANGLE_FAN").Int()
	UNSIGNED_BYTE = gl.Get("UNSIGNED_BYTE").Int()
	UNSIGNED_SHORT = gl.Get("UNSIGNED_SHORT").Int()
	VERTEX_SHADER = gl.Get("VERTEX_SHADER").Int()

	// buffers
	STATIC_DRAW = gl.Get("STATIC_DRAW").Int()
	STREAM_DRAW = gl.Get("STREAM_DRAW").Int()
	DYNAMIC_DRAW = gl.Get("DYNAMIC_DRAW").Int()
	ARRAY_BUFFER = gl.Get("ARRAY_BUFFER").Int()
	ELEMENT_ARRAY_BUFFER = gl.Get("ELEMENT_ARRAY_BUFFER").Int()
	BUFFER_SIZE = gl.Get("BUFFER_SIZE").Int()
	BUFFER_USAGE = gl.Get("BUFFER_USAGE").Int()

	// depth or stencil tests
	NEVER = gl.Get("NEVER").Int()
	LESS = gl.Get("LESS").Int()
	EQUAL = gl.Get("EQUAL").Int()
	LEQUAL = gl.Get("LEQUAL").Int()
	GREATER = gl.Get("GREATER").Int()
	NOTEQUAL = gl.Get("NOTEQUAL").Int()
	GEQUAL = gl.Get("GEQUAL").Int()
	ALWAYS = gl.Get("ALWAYS").Int()

	// framebuffers and renderbuffers
	FRAMEBUFFER = gl.Get("FRAMEBUFFER").Int()
	STENCIL_INDEX8 = gl.Get("STENCIL_INDEX8").Int()

	// stencil operations
	KEEP = gl.Get("KEEP").Int()
	REPLACE = gl.Get("REPLACE").Int()
	INCR = gl.Get("INCR").Int()
	DECR = gl.Get("DECR").Int()
	INVERT = gl.Get("INVERT").Int()
	INCR_WRAP = gl.Get("INCR_WRAP").Int()
	DECR_WRAP = gl.Get("DECR_WRAP").Int()

	// textures
	R8 = gl.Get("R8").Int()
	RGB = gl.Get("RGB").Int()
	RGBA = gl.Get("RGBA").Int()
	RGBA8 = gl.Get("RGBA8").Int()
}
