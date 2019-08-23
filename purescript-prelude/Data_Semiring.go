package purescript_prelude

import . "github.com/purescript-native/go-runtime"

func init() {
	exports := Foreign("Data.Semiring")

	exports["intAdd"] = func(x Any) Any {
		return func(y Any) Any {
			return x.(int) + y.(int)
		}
	}

	exports["intMul"] = func(x Any) Any {
		return func(y Any) Any {
			return x.(int) * y.(int)
		}
	}

	exports["numAdd"] = func(x Any) Any {
		return func(y Any) Any {
			return x.(float64) + y.(float64)
		}
	}

	exports["numMul"] = func(x Any) Any {
		return func(y Any) Any {
			return x.(float64) * y.(float64)
		}
	}

}
