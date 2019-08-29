package purescript_unsafe_coerce

import (
	. "github.com/purescript-native/go-runtime"
)

func init() {
	exports := Foreign("Unsafe.Coerce")

	exports["unsafeCoerce"] = func(x Any) Any {
		return x
	}
}
