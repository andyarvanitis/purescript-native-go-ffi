package purescript_assert

import (
	"Test_Assert"
	. "purescript"
)

func init() {
	exports := Test_Assert.Foreign

	exports["assert'"] = func(message Any) Any {
		return func(success Any) Any {
			return func() Any {
				if !(success.(bool)) {
					panic(message.(string))
				}
				return nil
			}
		}
	}

}
