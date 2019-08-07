package purescript_assert

import . "github.com/purescript-native/go-runtime"

func init() {
	exports := Foreign("Test.Assert")

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
