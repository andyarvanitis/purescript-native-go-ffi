package purescript_global

import (
	. "github.com/purescript-native/go-runtime"
	"math"
	"strconv"
)

func init() {
	exports := Foreign("Global")

	exports["infinity"] = math.Inf(1)

	exports["readFloat"] = func(x Any) Any {
		f, err := strconv.ParseFloat(x.(string), 64)
		if err == nil {
			return f
		}
		return math.NaN()
	}
	exports["isNaN"] = func(x Any) Any {
		return math.IsNaN(x.(float64))
	}

}
