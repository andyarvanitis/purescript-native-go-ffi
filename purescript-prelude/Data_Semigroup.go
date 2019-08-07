package purescript_prelude

import . "github.com/purescript-native/go-runtime"

func init() {
	exports := Foreign("Data.Semigroup")

	exports["concatString"] = func(s1 Any) Any {
		return func(s2 Any) Any {
			return s1.(string) + s2.(string)
		}
	}

	exports["concatArray"] = func(xs_ Any) Any {
		return func(ys_ Any) Any {
			xs, _ := xs_.([]Any)
			ys, _ := ys_.([]Any)
			result := make([]Any, 0, len(xs)+len(ys))
			result = append(result, xs...)
			result = append(result, ys...)
			return result
		}
	}

}
