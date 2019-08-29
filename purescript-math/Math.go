package purescript_math

import (
	. "github.com/purescript-native/go-runtime"
	"math"
)

func init() {
	exports := Foreign("Math")

	exports["abs"] = func(x_ Any) Any {
		x, _ := x_.(float64)
		return math.Abs(x)
	}

	exports["floor"] = func(x_ Any) Any {
		x, _ := x_.(float64)
		return math.Floor(x)
	}

	exports["pow"] = func(n_ Any) Any {
		return func(p_ Any) Any {
			n, _ := n_.(float64)
			p, _ := p_.(float64)
			return math.Pow(n, p)
		}
	}

	exports["remainder"] = func(n_ Any) Any {
		return func(m_ Any) Any {
			n, _ := n_.(float64)
			m, _ := m_.(float64)
			return math.Remainder(n, m)
		}
	}

}
