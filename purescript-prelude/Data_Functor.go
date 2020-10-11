package purescript_prelude

import . "github.com/purescript-native/go-runtime"

func init() {
	exports := Foreign("Data.Functor")

	exports["arrayMap"] = func(f Any) Any {
		return func(xs_ Any) Any {
			xs := xs_.([]Any)
			result := make([]Any, 0, len(xs))
			for _, elem := range xs {
				result = append(result, Apply(f, elem))
			}
			return result
		}
	}

}
