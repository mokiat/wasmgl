//go:build js && wasm

package wasmgl

import "syscall/js"

// Reference on WebGL2 constants:
// https://developer.mozilla.org/en-US/docs/Web/API/WebGL_API/Constants

// TODO: Hardcode the values of these constants instead of resolving them.
// This would allow them to be usable before the GL context is initialized.

var (
	NO_ERROR                    int
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
	REPEAT                      int
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
	UNSIGNED_BYTE               int
	UNSIGNED_SHORT              int
	VERTEX_SHADER               int

	// blending equations
	FUNC_ADD              int
	FUNC_SUBTRACT         int
	FUNC_REVERSE_SUBTRACT int

	// blending modes
	ZERO                     int
	ONE                      int
	SRC_COLOR                int
	ONE_MINUS_SRC_COLOR      int
	SRC_ALPHA                int
	ONE_MINUS_SRC_ALPHA      int
	DST_ALPHA                int
	ONE_MINUS_DST_ALPHA      int
	DST_COLOR                int
	ONE_MINUS_DST_COLOR      int
	SRC_ALPHA_SATURATE       int
	CONSTANT_COLOR           int
	ONE_MINUS_CONSTANT_COLOR int
	CONSTANT_ALPHA           int
	ONE_MINUS_CONSTANT_ALPHA int

	// buffers
	STATIC_DRAW          int
	STREAM_DRAW          int
	DYNAMIC_DRAW         int
	ARRAY_BUFFER         int
	ELEMENT_ARRAY_BUFFER int
	BUFFER_SIZE          int
	BUFFER_USAGE         int
	PIXEL_PACK_BUFFER    int

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
	FRAMEBUFFER              int
	STENCIL_INDEX8           int
	COLOR_ATTACHMENT0        int
	DEPTH_ATTACHMENT         int
	STENCIL_ATTACHMENT       int
	DEPTH_STENCIL_ATTACHMENT int
	FRAMEBUFFER_COMPLETE     int
	READ_FRAMEBUFFER         int
	DRAW_FRAMEBUFFER         int

	// pixel formats
	DEPTH_COMPONENT    int
	DEPTH_COMPONENT24  int
	DEPTH_COMPONENT32F int
	DEPTH24_STENCIL8   int

	// rendering primitives
	POINTS         int
	LINES          int
	LINE_LOOP      int
	LINE_STRIP     int
	TRIANGLES      int
	TRIANGLE_STRIP int
	TRIANGLE_FAN   int

	// stencil operations
	KEEP      int
	REPLACE   int
	INCR      int
	DECR      int
	INVERT    int
	INCR_WRAP int
	DECR_WRAP int

	// sync objects
	OBJECT_TYPE                int
	SYNC_CONDITION             int
	SYNC_STATUS                int
	SYNC_FLAGS                 int
	SYNC_FENCE                 int
	SYNC_GPU_COMMANDS_COMPLETE int
	UNSIGNALED                 int
	SIGNALED                   int
	ALREADY_SIGNALED           int
	TIMEOUT_EXPIRED            int
	CONDITION_SATISFIED        int
	WAIT_FAILED                int
	SYNC_FLUSH_COMMANDS_BIT    int

	// textures
	RED             int
	R8              int
	RGB             int
	RGBA            int
	RGBA8           int
	RGB16F          int
	RGBA16F         int
	RGB32F          int
	RGBA32F         int
	MIRRORED_REPEAT int
)

func initConstants(gl js.Value) {
	NO_ERROR = gl.Get("NO_ERROR").Int()
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
	REPEAT = gl.Get("REPEAT").Int()
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
	UNSIGNED_BYTE = gl.Get("UNSIGNED_BYTE").Int()
	UNSIGNED_SHORT = gl.Get("UNSIGNED_SHORT").Int()
	VERTEX_SHADER = gl.Get("VERTEX_SHADER").Int()

	// blending equations
	FUNC_ADD = gl.Get("FUNC_ADD").Int()
	FUNC_SUBTRACT = gl.Get("FUNC_SUBTRACT").Int()
	FUNC_REVERSE_SUBTRACT = gl.Get("FUNC_REVERSE_SUBTRACT").Int()

	// blending modes
	ZERO = gl.Get("ZERO").Int()
	ONE = gl.Get("ONE").Int()
	SRC_COLOR = gl.Get("SRC_COLOR").Int()
	ONE_MINUS_SRC_COLOR = gl.Get("ONE_MINUS_SRC_COLOR").Int()
	SRC_ALPHA = gl.Get("SRC_ALPHA").Int()
	ONE_MINUS_SRC_ALPHA = gl.Get("ONE_MINUS_SRC_ALPHA").Int()
	DST_ALPHA = gl.Get("DST_ALPHA").Int()
	ONE_MINUS_DST_ALPHA = gl.Get("ONE_MINUS_DST_ALPHA").Int()
	DST_COLOR = gl.Get("DST_COLOR").Int()
	ONE_MINUS_DST_COLOR = gl.Get("ONE_MINUS_DST_COLOR").Int()
	SRC_ALPHA_SATURATE = gl.Get("SRC_ALPHA_SATURATE").Int()
	CONSTANT_COLOR = gl.Get("CONSTANT_COLOR").Int()
	ONE_MINUS_CONSTANT_COLOR = gl.Get("ONE_MINUS_CONSTANT_COLOR").Int()
	CONSTANT_ALPHA = gl.Get("CONSTANT_ALPHA").Int()
	ONE_MINUS_CONSTANT_ALPHA = gl.Get("ONE_MINUS_CONSTANT_ALPHA").Int()

	// buffers
	STATIC_DRAW = gl.Get("STATIC_DRAW").Int()
	STREAM_DRAW = gl.Get("STREAM_DRAW").Int()
	DYNAMIC_DRAW = gl.Get("DYNAMIC_DRAW").Int()
	ARRAY_BUFFER = gl.Get("ARRAY_BUFFER").Int()
	ELEMENT_ARRAY_BUFFER = gl.Get("ELEMENT_ARRAY_BUFFER").Int()
	BUFFER_SIZE = gl.Get("BUFFER_SIZE").Int()
	BUFFER_USAGE = gl.Get("BUFFER_USAGE").Int()
	PIXEL_PACK_BUFFER = gl.Get("PIXEL_PACK_BUFFER").Int()

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
	COLOR_ATTACHMENT0 = gl.Get("COLOR_ATTACHMENT0").Int()
	DEPTH_ATTACHMENT = gl.Get("DEPTH_ATTACHMENT").Int()
	STENCIL_ATTACHMENT = gl.Get("STENCIL_ATTACHMENT").Int()
	DEPTH_STENCIL_ATTACHMENT = gl.Get("DEPTH_STENCIL_ATTACHMENT").Int()
	FRAMEBUFFER_COMPLETE = gl.Get("FRAMEBUFFER_COMPLETE").Int()
	READ_FRAMEBUFFER = gl.Get("READ_FRAMEBUFFER").Int()
	DRAW_FRAMEBUFFER = gl.Get("DRAW_FRAMEBUFFER").Int()

	// pixel formats
	DEPTH_COMPONENT = gl.Get("DEPTH_COMPONENT").Int()
	DEPTH_COMPONENT24 = gl.Get("DEPTH_COMPONENT24").Int()
	DEPTH_COMPONENT32F = gl.Get("DEPTH_COMPONENT32F").Int()
	DEPTH24_STENCIL8 = gl.Get("DEPTH24_STENCIL8").Int()

	// rendering primitives
	POINTS = gl.Get("POINTS").Int()
	LINES = gl.Get("LINES").Int()
	LINE_LOOP = gl.Get("LINE_LOOP").Int()
	LINE_STRIP = gl.Get("LINE_STRIP").Int()
	TRIANGLES = gl.Get("TRIANGLES").Int()
	TRIANGLE_STRIP = gl.Get("TRIANGLE_STRIP").Int()
	TRIANGLE_FAN = gl.Get("TRIANGLE_FAN").Int()

	// stencil operations
	KEEP = gl.Get("KEEP").Int()
	REPLACE = gl.Get("REPLACE").Int()
	INCR = gl.Get("INCR").Int()
	DECR = gl.Get("DECR").Int()
	INVERT = gl.Get("INVERT").Int()
	INCR_WRAP = gl.Get("INCR_WRAP").Int()
	DECR_WRAP = gl.Get("DECR_WRAP").Int()

	// sync objects
	OBJECT_TYPE = gl.Get("OBJECT_TYPE").Int()
	SYNC_CONDITION = gl.Get("SYNC_CONDITION").Int()
	SYNC_STATUS = gl.Get("SYNC_STATUS").Int()
	SYNC_FLAGS = gl.Get("SYNC_FLAGS").Int()
	SYNC_FENCE = gl.Get("SYNC_FENCE").Int()
	SYNC_GPU_COMMANDS_COMPLETE = gl.Get("SYNC_GPU_COMMANDS_COMPLETE").Int()
	UNSIGNALED = gl.Get("UNSIGNALED").Int()
	SIGNALED = gl.Get("SIGNALED").Int()
	ALREADY_SIGNALED = gl.Get("ALREADY_SIGNALED").Int()
	TIMEOUT_EXPIRED = gl.Get("TIMEOUT_EXPIRED").Int()
	CONDITION_SATISFIED = gl.Get("CONDITION_SATISFIED").Int()
	WAIT_FAILED = gl.Get("WAIT_FAILED").Int()
	SYNC_FLUSH_COMMANDS_BIT = gl.Get("SYNC_FLUSH_COMMANDS_BIT").Int()

	// textures
	RED = gl.Get("RED").Int()
	R8 = gl.Get("R8").Int()
	RGB = gl.Get("RGB").Int()
	RGBA = gl.Get("RGBA").Int()
	RGBA8 = gl.Get("RGBA8").Int()
	RGB16F = gl.Get("RGB16F").Int()
	RGBA16F = gl.Get("RGBA16F").Int()
	RGB32F = gl.Get("RGB32F").Int()
	RGBA32F = gl.Get("RGBA32F").Int()
	MIRRORED_REPEAT = gl.Get("MIRRORED_REPEAT").Int()
}
