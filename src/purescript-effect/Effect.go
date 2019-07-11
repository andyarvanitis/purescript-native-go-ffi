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

	exports["bindE"] = func(a Any) Any {
		return func(f Any) Any {
			return func() Any {
				return Run(Apply(f, Run(a)))
			}
		}
	}

}
