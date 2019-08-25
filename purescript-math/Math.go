package purescript_math

import (
	. "github.com/purescript-native/go-runtime"
	"math"
)

func init() {
	exports := Foreign("Math")

	exports["floor"] = func(x_ Any) Any {
		x, _ := x_.(float64)
		return math.Floor(x)
	}

	exports["remainder"] = func(n_ Any) Any {
		return func(m_ Any) Any {
			n, _ := n_.(float64)
			m, _ := m_.(float64)
			return math.Remainder(n, m)
		}
	}

}
