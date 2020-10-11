package purescript_unfoldable

import . "github.com/purescript-native/go-runtime"

func init() {
	exports := Foreign("Data.Unfoldable")

	exports["unfoldrArrayImpl"] = func(isNothing_ Any) Any {
		return func(fromJust_ Any) Any {
			return func(fst_ Any) Any {
				return func(snd_ Any) Any {
					return func(f_ Any) Any {
						return func(b Any) Any {
							isNothing := isNothing_.(Fn)
							fromJust := fromJust_.(Fn)
							fst := fst_.(Fn)
							snd := snd_.(Fn)
							f := f_.(Fn)
							result := make([]Any, 0)
							value := b
							for {
								maybe := f(value)
								nothing := isNothing(maybe).(bool)
								if nothing {
									return result
								}
								tuple := fromJust(maybe)
								result = append(result, fst(tuple))
								value = snd(tuple)
							}
						}
					}
				}
			}
		}
	}

}
