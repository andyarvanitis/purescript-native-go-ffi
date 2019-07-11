package purescript_prelude

import (
	"Data_Semiring"
	. "purescript"
)

func init() {
	exports := Data_Semiring.Foreign

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
