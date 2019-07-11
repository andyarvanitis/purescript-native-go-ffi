package purescript_prelude

import (
	"Data_Semigroup"
	. "purescript"
)

func init() {
	exports := Data_Semigroup.Foreign

	exports["concatString"] = func(s1 Any) Any {
		return func(s2 Any) Any {
			return s1.(string) + s2.(string)
		}
	}

	exports["concatArray"] = func(xs_ Any) Any {
		return func(ys_ Any) Any {
			xs, _ := xs_.(Array)
			ys, _ := ys_.(Array)
			result := make(Array, 0, len(xs)+len(ys))
			result = append(result, xs...)
			result = append(result, ys...)
			return result
		}
	}

}
