package purescript_effect

import (
	"Effect_"
	. "purescript"
)

func init() {
	exports := Effect_.Foreign

	exports["pureE"] = func(a Any) Any {
		return func() Any {
			return a
		}
	}

}
