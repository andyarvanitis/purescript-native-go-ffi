package purescript_effect

import (
	. "github.com/purescript-native/go-runtime"
)

func init() {
	exports := Foreign("Effect.Unsafe")

	exports["unsafePerformEffect"] = func(f Any) Any {
		return Run(f)
	}
}
