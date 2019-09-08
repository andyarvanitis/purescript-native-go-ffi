package purescript_unsafe_coerce

import (
	. "github.com/purescript-native/go-runtime"
)

func init() {
	exports := Foreign("Unsafe.Coerce")

	exports["unsafeCoerce"] = func(x_ Any) Any {
		// Temporary solution for the usage of copyImpl in both `unsafeFreeze` and `unsafeThaw` of STArray
		switch x := x_.(type) {
		case []Any:
			return &x
		case *[]Any:
			return *x
		default:
			return x
		}
	}
}
