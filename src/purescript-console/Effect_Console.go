package console

import "fmt"
import . "purescript"
import "Effect_Console"

type Loader = Any

func init() {
	exports := Effect_Console.Foreign

	exports["log"] = func(s Any) Any {
		return func() Any {
			fmt.Println(s.(string))
			return nil
		}
	}
}
