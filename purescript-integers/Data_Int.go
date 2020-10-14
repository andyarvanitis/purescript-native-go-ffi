package purescript_integers

import (
	. "github.com/purescript-native/go-runtime"
	"math"
	"strconv"
)

func init() {
	exports := Foreign("Data.Int")

	exports["toNumber"] = func(n Any) Any {
		return float64(n.(int))
	}

	exports["fromNumberImpl"] = func(just_ Any) Any {
		return func(nothing Any) Any {
			return func(n_ Any) Any {
				just := just_.(Fn)
				n := n_.(float64)
				if math.Trunc(n) == n {
					return just(int(n))
				} else {
					return nothing
				}
			}
		}
	}

	exports["fromStringAsImpl"] = func(just_ Any) Any {
		return func(nothing Any) Any {
			return func(radix Any) Any {
				return func(s Any) Any {
					just := just_.(Fn)
					i, err := strconv.ParseInt(s.(string), radix.(int), 32)
					if err == nil {
						return just(int(i))
					} else {
						return nothing
					}
				}
			}
		}
	}

	exports["quot"] = func(x Any) Any {
		return func(y Any) Any {
			return x.(int) / y.(int)
		}
	}

	exports["rem"] = func(x Any) Any {
		return func(y Any) Any {
			return x.(int) % y.(int)
		}
	}

	exports["pow"] = func(n_ Any) Any {
		return func(p_ Any) Any {
			n := n_.(int)
			p := p_.(int)
			return int(math.Pow(float64(n), float64(p)))
		}
	}

	exports["toStringAs"] = func(radix Any) Any {
		return func(i Any) Any {
			return strconv.FormatInt(int64(i.(int)), radix.(int))
		}
	}

}
