package purescript_prelude

import . "github.com/purescript-native/go-runtime"

func init() {

	exports := Foreign("Control.Bind")

	exports["arrayBind"] = func(xs_ Any) Any {
		return func(f_ Any) Any {
			xs, f := xs_.([]Any), f_.(Fn)
			result := []Any{}
			for _, x := range xs {
				result = append(result, f(x).([]Any)...)
			}
			return result
		}
	}

}
