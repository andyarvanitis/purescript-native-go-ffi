package purescript_foldable_traversable

import (
	. "github.com/purescript-native/go-runtime"
)

func init() {
	exports := Foreign("Data.FunctorWithIndex")

	exports["mapWithIndexArray"] = func(f Any) Any {
		return func(xs_ Any) Any {
			xs, _ := xs_.([]Any)
			result := make([]Any, 0, len(xs))
			for i, x := range xs {
				result = append(result, Apply(f, i, x))
			}
			return result
		}
	}

}
