package purescript_arrays

import . "purescript"
import "Data_Array"

func init() {
	exports := Data_Array.Foreign

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
			ns := make([]Any, length)
			for i := start; i != end; i += step {
				ns = append(ns, i)
			}
			return append(ns, end)
		}
	}

	exports["length"] = func(xs Any) Any {
		return len(xs.([]Any))
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
						return just.(Fn)(xs[i])
					}
				}
			}
		}
	}
}
