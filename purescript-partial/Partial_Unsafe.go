package purescript_partial

import . "github.com/purescript-native/go-runtime"

func init() {
	exports := Foreign("Partial.Unsafe")

	exports["unsafePartial"] = func(f Any) Any {
		return Apply(f, nil)
	}

}
