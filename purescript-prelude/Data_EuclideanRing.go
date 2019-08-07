package purescript_prelude

import (
	. "github.com/purescript-native/go-runtime"
	"math"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func init() {
	exports := Foreign("Data.EuclideanRing")

	exports["intDegree"] = func(x_ Any) Any {
		x, _ := x_.(int)
		return min(abs(x), maxInt)
	}

	// See the Euclidean definition in
	// https://en.m.wikipedia.org/wiki/Modulo_operation.
	//
	exports["intDiv"] = func(x_ Any) Any {
		return func(y_ Any) Any {
			x, _ := x_.(int)
			y, _ := y_.(int)
			if y == 0 {
				return 0
			} else {
				xf := float64(x)
				yf := float64(y)
				if y > 0 {
					return int(math.Floor(xf / yf))
				} else {
					return -int(math.Floor(xf / -yf))
				}
			}
		}
	}

	exports["intMod"] = func(x_ Any) Any {
		return func(y_ Any) Any {
			x, _ := x_.(int)
			y, _ := y_.(int)
			if y == 0 {
				return 0
			} else {
				yy := abs(y)
				return ((x % yy) + yy) % yy
			}
		}
	}

	exports["numDiv"] = func(n1 Any) Any {
		return func(n2 Any) Any {
			return n1.(float64) / n2.(float64)
		}
	}

}
