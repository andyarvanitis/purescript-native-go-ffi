package purescript_nullable

import (
	. "github.com/purescript-native/go-runtime"
)

func init() {
	exports := Foreign("Data.Nullable")

	exports["null"] = nil

	exports["nullable"] = func(a, r, f Any) Any {
		if a == nil {
			return r
		}
		return Apply(f, a)
	}

	exports["notNull"] = func(x Any) Any {
		return x
	}
}
