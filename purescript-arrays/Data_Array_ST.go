package purescript_arrays

import (
	"fmt"

	. "github.com/purescript-native/go-runtime"
)

func init() {

	exports := Foreign("Data.Array.ST")

	exports["empty"] = func() Any {
		return new([]Any)
	}

	exports["peekImpl"] = func(just_ Any) Any {
		return func(nothing Any) Any {
			return func(i_ Any) Any {
				return func(xs_ Any) Any {
					return func() Any {
						just, i, xs := just_.(Fn), i_.(int), xs_.(*[]Any)
						if i >= 0 && i < len(*xs) {
							return just((*xs)[i])
						}
						return nothing
					}
				}
			}
		}
	}

	exports["poke"] = func(i_ Any) Any {
		return func(a Any) Any {
			return func(xs_ Any) Any {
				return func() Any {
					i, xs := i_.(int), xs_.(*[]Any)
					result := i >= 0 && i < len(*xs)
					if result {
						(*xs)[i] = a
					}
					return result
				}
			}
		}
	}

	exports["pushAll"] = func(as_ Any) Any {
		return func(xs_ Any) Any {
			return func() Any {
				as := as_.([]Any)
				xs := xs_.(*[]Any)
				*xs = append(*xs, as...)
				return len(*xs)
			}
		}
	}

	exports["copyImpl"] = func(xs_ Any) Any {
		return func() Any {
			// Temporary solution for the usage of copyImpl in both `freeze` and `thaw` of STArray
			switch xs := xs_.(type) {
			case []Any:
				fmt.Println(xs)
				ys := append([]Any{}, xs...)
				return &ys
			case *[]Any:
				fmt.Println(*xs)
				return append([]Any{}, *xs...)
			default:
				panic("The value passed to copyImpl is neither an Array nor Array pointer")
			}
		}
	}

}
