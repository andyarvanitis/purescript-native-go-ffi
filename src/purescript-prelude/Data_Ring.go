package purescript_prelude

import (
	. "purescript"
)

func init() {
	exports := Foreign("Data.Ring")

	exports["intSub"] = func(x Any) Any {
		return func(y Any) Any {
			return x.(int) - y.(int)
		}
	}

	exports["numSub"] = func(x Any) Any {
		return func(y Any) Any {
			return x.(float64) - y.(float64)
		}
	}
}
