package purescript_foldable_traversable

import . "github.com/purescript-native/go-runtime"

func init() {
	exports := Foreign("Data.Foldable")

	exports["foldrArray"] = func(f Any) Any {
		return func(init Any) Any {
			return func(xs_ Any) Any {
				xs, _ := xs_.([]Any)
				var acc = init
				length := len(xs)
				for i := length - 1; i >= 0; i-- {
					acc = Apply(f, xs[i], acc)
				}
				return acc
			}
		}
	}

	exports["foldlArray"] = func(f Any) Any {
		return func(init Any) Any {
			return func(xs_ Any) Any {
				xs, _ := xs_.([]Any)
				var acc = init
				for _, x := range xs {
					acc = Apply(f, acc, x)
				}

				return acc
			}
		}
	}

}
