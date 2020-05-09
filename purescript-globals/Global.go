package purescript_global

import (
	. "github.com/purescript-native/go-runtime"
	"math"
	"strconv"
)

func init() {
	exports := Foreign("Global")

	exports["infinity"] = math.Inf(1)

	exports["readFloat"] = func(x_ Any) Any {
		x, _ := x_.(string)
		f, err := strconv.ParseFloat(x, 64)
		if err == nil {
			return f
		}
		return math.NaN()
	}
	exports["isNaN"] = func(x_ Any) Any {
		x, _ := x_.(float64)
		return math.IsNaN(x)
	}

}
