package purescript_console

import (
	"fmt"
	. "purescript"
)

func init() {
	exports := Foreign("Effect.Console")

	exports["log"] = func(s Any) Any {
		return func() Any {
			fmt.Println(s)
			return nil
		}
	}

	exports["warn"] = func(s Any) Any {
		return func() Any {
			fmt.Println(s)
			return nil
		}
	}
}
