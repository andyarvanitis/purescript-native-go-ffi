package purescript_arrays

import . "github.com/purescript-native/go-runtime"

func init() {

	exports := Foreign("Data.Array.ST")

	exports["empty"] = func() Any {
		return []Any{}
	}

	exports["pushAll"] = func(as_ Any) Any {
		return func(xs_ Any) Any {
			return func() Any {
				// TODO: buggy implementation, we need to update the reference not create a new slice
				as := (as_).([]Any)
				xs := xs_.([]Any)
				as = append(as, xs...)
				return len(as)
			}
		}
	}

	exports["copyImpl"] = func(xs_ Any) Any {
		return func() Any {
			xs := xs_.([]Any)
			ys := append([]Any{}, xs...)
			return ys
		}
	}

}
