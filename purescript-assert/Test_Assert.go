package purescript_assert

import (
	"errors"
	. "github.com/purescript-native/go-runtime"
)

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

	exports["checkThrows"] = func(f Any) Any {
		return func() Any {
			result := false
			func() {
				defer func() {
					if e := recover(); e != nil {
						switch e.(type) {
						case error:
							result = true
						default:
							err := errors.New("Threw something other than an 'error'")
							panic(err)
						}
					}
				}()
				Run(f)
			}()
			return result
		}
	}

}
