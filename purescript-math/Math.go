package purescript_math

import (
	. "github.com/purescript-native/go-runtime"
	"math"
)

func init() {
	exports := Foreign("Math")

	exports["abs"] = func(x Any) Any {
		return math.Abs(x.(float64))
	}

	exports["floor"] = func(x Any) Any {
		return math.Floor(x.(float64))
	}

	exports["pow"] = func(n Any) Any {
		return func(p Any) Any {
			return math.Pow(n.(float64), p.(float64))
		}
	}

	exports["remainder"] = func(n Any) Any {
		return func(m Any) Any {
			return math.Remainder(n.(float64), m.(float64))
		}
	}

	exports["round"] = func(x Any) Any {
		return math.Round(x.(float64))
	}
}
