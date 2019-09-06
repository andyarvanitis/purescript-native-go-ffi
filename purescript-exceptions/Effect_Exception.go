package purescript_exceptions

import (
	"errors"
	"fmt"

	. "github.com/purescript-native/go-runtime"
)

func init() {
	exports := Foreign("Effect.Exception")

	exports["error"] = func(msg Any) Any {
		return errors.New(msg.(string))
	}

	exports["message"] = func(e_ Any) Any {
		e, _ := e_.(error)
		return e.Error()
	}

	exports["throwException"] = func(e Any) Any {
		return func() Any {
			panic(e)
		}
	}

	exports["catchException"] = func(c_ Any) Any {
		return func(t_ Any) Any {
			return func() Any {
				c, _ := c_.(Fn)
				t, _ := t_.(EffFn)
				var result Any = nil
				func() {
					defer func() {
						if e := recover(); e != nil {
							switch e.(type) {
							case error:
								result = Run(c(e))
							default:
								result = Run(c(errors.New(fmt.Sprint(e))))
							}
						}
					}()
					result = t()
				}()
				return result
			}
		}
	}

}
