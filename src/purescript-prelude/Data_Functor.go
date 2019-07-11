package purescript_prelude

import (
	"Data_Functor"
	. "purescript"
)

func init() {
	exports := Data_Functor.Foreign

	exports["arrayMap"] = func(f Any) Any {
		return func(xs_ Any) Any {
			xs, _ := xs_.(Array)
			result := make(Array, 0, len(xs))
			for i, elem := range xs {
				result[i] = Apply(f, elem)
			}
			return result
		}
	}

}
