package purescript_arrays

import . "github.com/purescript-native/go-runtime"

func init() {

	exports := Foreign("Data.Array")

	//------------------------------------------------------------------------------
	// Array creation --------------------------------------------------------------
	//------------------------------------------------------------------------------

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

	exports["replicate"] = func(count_ Any) Any {
		return func(value Any) Any {
			var count = count_.(int)
			if count < 1 {
				return []Any{}
			}
			var arr = make([]Any, count)
			for i := range arr {
				arr[i] = value
			}
			return arr
		}
	}

	exports["fromFoldableImpl"] = func(foldr Any) Any {
		return func(foldable Any) Any {
			var f = func(x Any) Any {
				return func(acc Any) Any {
					return append(acc.([]Any), x)
				}
			}
			var xs = Apply(foldr, f, []Any{}, foldable)
			return exports["reverse"].(Fn)(xs)
		}
	}

	//------------------------------------------------------------------------------
	// Array size ------------------------------------------------------------------
	//------------------------------------------------------------------------------

	exports["length"] = func(xs_ Any) Any {
		xs, _ := xs_.([]Any)
		return len(xs)
	}

	//------------------------------------------------------------------------------
	// Extending arrays ------------------------------------------------------------
	//------------------------------------------------------------------------------

	//------------------------------------------------------------------------------
	// Non-indexed reads -----------------------------------------------------------
	//------------------------------------------------------------------------------

	//------------------------------------------------------------------------------
	// Indexed operations ----------------------------------------------------------
	//------------------------------------------------------------------------------

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

	//------------------------------------------------------------------------------
	// Transformations -------------------------------------------------------------
	//------------------------------------------------------------------------------

	exports["reverse"] = func(xs_ Any) Any {
		xs := xs_.([]Any)
		l := len(xs)
		ys := make([]Any, l)
		for i, j := 0, l-1; i < l; i, j = i+1, j-1 {
			ys[i] = xs[j]
		}
		return ys
	}

	//------------------------------------------------------------------------------
	// Sorting ---------------------------------------------------------------------
	//------------------------------------------------------------------------------

	//------------------------------------------------------------------------------
	// Subarrays -------------------------------------------------------------------
	//------------------------------------------------------------------------------

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

	//------------------------------------------------------------------------------
	// Zipping ---------------------------------------------------------------------
	//------------------------------------------------------------------------------

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
