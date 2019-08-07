package purescript_prelude

import . "github.com/purescript-native/go-runtime"

func init() {
	exports := Foreign("Data.HeytingAlgebra")

	exports["boolConj"] = func(x Any) Any {
		return func(y Any) Any {
			return x.(bool) && y.(bool)
		}
	}

	exports["boolDisj"] = func(x Any) Any {
		return func(y Any) Any {
			return x.(bool) || y.(bool)
		}
	}

	exports["boolNot"] = func(b Any) Any {
		return !b.(bool)
	}

}
