package purescript_partial

import . "purescript"

func init() {
	exports := Foreign("Partial.Unsafe")

	exports["unsafePartial"] = func(f Any) Any {
		return Apply(f, nil)
	}

}
