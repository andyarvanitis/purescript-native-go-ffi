package purescript_arrays

import (
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

	exports["unsafeFreeze"] = func(xs_ Any) Any {
		return func() Any {
			xs, _ := xs_.(*[]Any)
			return *xs
		}
	}

	exports["unsafeThaw"] = func(xs_ Any) Any {
		return func() Any {
			xs, _ := xs_.([]Any)
			return &xs
		}
	}

	exports["freeze"] = func(xs_ Any) Any {
		return func() Any {
			xs, _ := xs_.(*[]Any)
			return append([]Any{}, *xs...)
		}
	}

	exports["thaw"] = func(xs_ Any) Any {
		return func() Any {
			xs, _ := xs_.([]Any)
			result := append([]Any{}, xs...)
			return &result
		}
	}

}
