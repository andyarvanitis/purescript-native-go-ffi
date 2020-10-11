package purescript_integers

import (
	"fmt"
	. "github.com/purescript-native/go-runtime"
	"math"
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
				if math.Round(n) == n {
					return just(int(n))
				} else {
					return nothing
				}
			}
		}
	}

	exports["pow"] = func(n_ Any) Any {
		return func(p_ Any) Any {
			n := n_.(int)
			p := p_.(int)
			return int(math.Pow(float64(n), float64(p)))
		}
	}

	exports["toStringAs"] = func(radix_ Any) Any {
		return func(i_ Any) Any {
			radix := radix_.(int)
			i := i_.(int)
			switch radix {
			case 2:
				return fmt.Sprintf("%b", i)
			case 8:
				return fmt.Sprintf("%o", i)
			case 10:
				return fmt.Sprintf("%d", i)
			case 16:
				return fmt.Sprintf("%x", i)
			default:
				panic(fmt.Sprintf("Unsupported radix (%d)", radix))
			}
		}
	}

}
