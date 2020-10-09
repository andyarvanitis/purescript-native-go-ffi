package purescript_integers

import (
	. "github.com/purescript-native/go-runtime"
)

func init() {
	exports := Foreign("Data.Int.Bits")

	exports["zshr"] = func(n Any) Any {
		return func (i Any) Any {
			return int(uint(n.(int)) >> i.(int))
		}
	}
}
