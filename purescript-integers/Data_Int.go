package purescript_integers

import (
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
				just, _ := just_.(Fn)
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
			n, _ := n_.(int)
			p, _ := p_.(int)
			return int(math.Pow(float64(n), float64(p)))
		}
	}

}
