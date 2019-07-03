package console

import "fmt"
import . "purescript"
import "Effect_Console"

const Loader = true

func init() {
	exports := Effect_Console.Foreign

	exports["log"] = func(s Any) Any {
		return func() Any {
			fmt.Println(s.(string))
			return nil
		}
	}
}
