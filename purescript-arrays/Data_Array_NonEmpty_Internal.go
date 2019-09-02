package purescript_arrays

import . "github.com/purescript-native/go-runtime"

func init() {

	exports := Foreign("Data.Array.NonEmpty.Internal")

	exports["fold1Impl"] = func(f Any) Any {
		return func(xs_ Any) Any {
			xs, _ := xs_.([]Any)
			acc := xs[0]
			for i := 1; i < len(xs); i++ {
				acc = Apply(f, acc, xs[i])
			}
			return acc
		}
	}

}
