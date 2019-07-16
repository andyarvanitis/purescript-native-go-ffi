package purescript_prelude

import . "purescript"

func init() {
	exports := Foreign("Data.Functor")

	exports["arrayMap"] = func(f Any) Any {
		return func(xs_ Any) Any {
			xs, _ := xs_.(Array)
			result := make(Array, 0, len(xs))
			for _, elem := range xs {
				result = append(result, Apply(f, elem))
			}
			return result
		}
	}

}
