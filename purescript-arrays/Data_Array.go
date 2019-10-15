package purescript_arrays

import (
	"sort"

	. "github.com/purescript-native/go-runtime"
)

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

	exports["cons"] = func(e Any) Any {
		return func(l Any) Any {
			return append([]Any{e}, l.([]Any)...)
		}
	}

	exports["snoc"] = func(l_ Any) Any {
		return func(e Any) Any {
			l := l_.([]Any)
			xs := make([]Any, len(l), len(l)+1)
			copy(xs, l)
			return append(xs, e)
		}
	}

	//------------------------------------------------------------------------------
	// Non-indexed reads -----------------------------------------------------------
	//------------------------------------------------------------------------------

	exports["uncons'"] = func(empty_ Any) Any {
		return func(next Any) Any {
			return func(xs_ Any) Any {
				empty, xs := empty_.(Fn), xs_.([]Any)
				if len(xs) == 0 {
					return empty(Dict{})
				}
				return Apply(next, xs[0], xs[1:])
			}
		}
	}

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
					}
					return Apply(just, xs[i])
				}
			}
		}
	}

	exports["findIndexImpl"] = func(just_ Any) Any {
		return func(nothing Any) Any {
			return func(f_ Any) Any {
				return func(xs_ Any) Any {
					xs, f, just := xs_.([]Any), f_.(Fn), just_.(Fn)
					for i, x := range xs {
						if f(x).(bool) {
							return just(i)
						}
					}
					return nothing
				}
			}
		}
	}

	exports["findLastIndexImpl"] = func(just_ Any) Any {
		return func(nothing Any) Any {
			return func(f_ Any) Any {
				return func(xs_ Any) Any {
					xs, f, just := xs_.([]Any), f_.(Fn), just_.(Fn)
					for i := len(xs) - 1; i >= 0; i-- {
						if f(xs[i]).(bool) {
							return just(i)
						}
					}
					return nothing
				}
			}
		}
	}

	exports["_insertAt"] = func(just_ Any) Any {
		return func(nothing Any) Any {
			return func(i_ Any) Any {
				return func(a Any) Any {
					return func(xs_ Any) Any {
						just, xs, i := just_.(Fn), xs_.([]Any), i_.(int)
						if i < 0 || i > len(xs) {
							return nothing
						}
						ys := make([]Any, len(xs)+1)
						copy(ys, xs[:i])
						ys[i] = a
						copy(ys[i+1:], xs[i:])
						return just(ys)
					}
				}
			}
		}
	}

	exports["_deleteAt"] = func(just_ Any) Any {
		return func(nothing Any) Any {
			return func(i_ Any) Any {
				return func(xs_ Any) Any {
					just, i, xs := just_.(Fn), i_.(int), xs_.([]Any)
					if i < 0 || i >= len(xs) {
						return nothing
					}
					ys := make([]Any, len(xs)-1)
					copy(ys, xs[:i])
					copy(ys[i:], xs[i+1:])
					return just(ys)
				}
			}
		}
	}

	exports["_updateAt"] = func(just_ Any) Any {
		return func(nothing Any) Any {
			return func(i_ Any) Any {
				return func(a Any) Any {
					return func(xs_ Any) Any {
						just, i, xs := just_.(Fn), i_.(int), xs_.([]Any)
						if i < 0 || i >= len(xs) {
							return nothing
						}
						ys := make([]Any, len(xs))
						copy(ys, xs)
						ys[i] = a
						return just(ys)
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

	exports["concat"] = func(xss_ Any) Any {
		xss := xss_.([]Any)
		result := []Any{}
		for _, xs := range xss {
			result = append(result, xs.([]Any)...)
		}
		return result
	}

	exports["filter"] = func(f_ Any) Any {
		return func(xs_ Any) Any {
			xs, f := xs_.([]Any), f_.(Fn)
			result := []Any{}
			for _, x := range xs {
				if f(x).(bool) {
					result = append(result, x)
				}
			}
			return result
		}
	}

	exports["partition"] = func(f_ Any) Any {
		return func(xs_ Any) Any {
			xs, f := xs_.([]Any), f_.(Fn)
			result := Dict{"yes": []Any{}, "no": []Any{}}
			for _, x := range xs {
				if f(x).(bool) {
					result["yes"] = append(result["yes"].([]Any), x)
				} else {
					result["no"] = append(result["no"].([]Any), x)
				}
			}
			return result
		}
	}

	//------------------------------------------------------------------------------
	// Sorting ---------------------------------------------------------------------
	//------------------------------------------------------------------------------

	exports["sortImpl"] = func(f Any) Any {
		return func(l_ Any) Any {
			l := l_.([]Any)
			xs := make([]Any, len(l))
			copy(xs, l)
			sort.SliceStable(xs, func(i int, j int) bool {
				return Apply(f, xs[i], xs[j]).(int) < 0
			})
			return xs
		}
	}

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

	exports["take"] = func(n_ Any) Any {
		return func(l_ Any) Any {
			n, l := n_.(int), l_.([]Any)
			if n < 1 {
				return []Any{}
			}
			if n > len(l) {
				return l
			}
			return l[:n]
		}
	}

	exports["drop"] = func(n_ Any) Any {
		return func(l_ Any) Any {
			n, l := n_.(int), l_.([]Any)
			if n < 1 {
				return l
			}
			if n > len(l) {
				return []Any{}
			}
			return l[n:]
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

	//------------------------------------------------------------------------------
	// Partial ---------------------------------------------------------------------
	//------------------------------------------------------------------------------

	exports["unsafeIndexImpl"] = func(xs_ Any) Any {
		return func(n_ Any) Any {
			xs, n := xs_.([]Any), n_.(int)
			return xs[n]
		}
	}

}
