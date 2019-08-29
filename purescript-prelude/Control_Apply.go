package purescript_prelude

import . "github.com/purescript-native/go-runtime"

func init() {

	exports := Foreign("Control.Apply")

	exports["arrayApply"] = func(fs_ Any) Any {
		return func(xs_ Any) Any {
			fs, _ := fs_.([]Any)
			xs, _ := xs_.([]Any)
			result := make([]Any, 0, len(fs)*len(xs))
			for _, f := range fs {
				for _, x := range xs {
					result = append(result, Apply(f, x))
				}
			}
			return result
		}
	}

}
