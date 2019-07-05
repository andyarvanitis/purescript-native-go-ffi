package purescript_prelude

import (
	"Data_Show"
	"fmt"
	. "purescript"
)

func init() {
	exports := Data_Show.Foreign

	exports["showIntImpl"] = func(n Any) Any {
		return fmt.Sprint(n)
	}

	exports["showNumberImpl"] = func(n Any) Any {
		return fmt.Sprint(n)
	}

	exports["showStringImpl"] = func(s Any) Any {
		return "\"" + fmt.Sprint(s) + "\""
	}

	exports["showArrayImpl"] = func(f_ Any) Any {
		return func(xs_ Any) Any {
			f, _ := f_.(Fn)
			xs, _ := xs_.([]Any)
			result := "["
			length := len(xs)
			for count, val := range xs {
				result += fmt.Sprint(f(val))
				if count < length-1 {
					result += ","
				}
			}
			result += "]"
			return result
		}
	}

}
