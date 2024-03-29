package purescript_prelude

import (
	"fmt"
	"math"
	"strings"

	. "github.com/purescript-native/go-runtime"
)

func init() {
	exports := Foreign("Data.Show")

	exports["showIntImpl"] = func(n Any) Any {
		return fmt.Sprint(n)
	}

	exports["showNumberImpl"] = func(n Any) Any {
		f := n.(float64)
		s := fmt.Sprintf("%g", f)
		if strings.Contains(s, ".") || strings.Contains(s, "e") || math.IsNaN(f) {
			return s
		} else {
			return s + ".0"
		}
	}

	exports["showCharImpl"] = func(s Any) Any {
		return fmt.Sprintf("%q", s)
	}

	exports["showStringImpl"] = func(s Any) Any {
		return "\"" + fmt.Sprint(s) + "\""
	}

	exports["showArrayImpl"] = func(f Any) Any {
		return func(xs_ Any) Any {
			xs, _ := xs_.([]Any)
			result := "["
			length := len(xs)
			for count, val := range xs {
				result += fmt.Sprint(Apply(f, val))
				if count < length-1 {
					result += ","
				}
			}
			result += "]"
			return result
		}
	}

	exports["cons"] = func(head Any) Any {
		return func(tail_ Any) Any {
			return append([]Any{head}, tail_.([]Any)...)
		}
	}

	exports["join"] = func(separator Any) Any {
		return func(xs_ Any) Any {
			xs := xs_.([]Any)
			ss := make([]string, len(xs))
			for i, x := range xs {
				ss[i] = x.(string)
			}

			res := strings.Join(ss, separator.(string))
			return res
		}
	}

}
