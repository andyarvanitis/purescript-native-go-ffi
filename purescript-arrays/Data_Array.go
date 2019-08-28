package purescript_arrays

import . "github.com/purescript-native/go-runtime"

func init() {

	exports := Foreign("Data.Array")

	exports["range"] = func(start_ Any) Any {
		return func(end_ Any) Any {
			start, _ := start_.(int)
			end, _ := end_.(int)
			var length int
			var step int
			if start > end {
				length = start - end + 1
				step = -1
			} else {
				length = end - start + 1
				step = 1
			}
			ns := make([]Any, 0, length)
			for i := start; i != end; i += step {
				ns = append(ns, i)
			}
			return append(ns, end)
		}
	}

	exports["length"] = func(xs_ Any) Any {
		xs, _ := xs_.([]Any)
		return len(xs)
	}

	exports["indexImpl"] = func(just Any) Any {
		return func(nothing Any) Any {
			return func(xs_ Any) Any {
				return func(i_ Any) Any {
					xs, _ := xs_.([]Any)
					i, _ := i_.(int)
					if i < 0 || i >= len(xs) {
						return nothing
					} else {
						return Apply(just, xs[i])
					}
				}
			}
		}
	}

	exports["slice"] = func(s_ Any) Any {
		return func(e_ Any) Any {
			return func(l_ Any) Any {
				s := s_.(int)
				e := e_.(int)
				l := l_.([]Any)
				sz := len(l)
				if s < 0 {
					s = sz + s
				}
				if e < 0 {
					e = sz + e
				}
				return l[s:e]
			}
		}
	}

	exports["zipWith"] = func(f_ Any) Any {
		return func(xs_ Any) Any {
			return func(ys_ Any) Any {
				f, _ := f_.(Fn)
				xs, _ := xs_.([]Any)
				ys, _ := ys_.([]Any)
				lxs := len(xs)
				l := len(ys)
				if lxs < l {
					l = lxs
				}
				result := make([]Any, 0, l)
				for i := 0; i < l; i++ {
					fx, _ := f(xs[i]).(Fn)
					result = append(result, fx(ys[i]))
				}
				return result
			}
		}
	}

}
