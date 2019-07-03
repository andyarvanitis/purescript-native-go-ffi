package prelude

import "strconv"
import . "purescript"
import "Data_Show"

type Loader = Any

func init() {
	exports := Data_Show.Foreign

	exports["showIntImpl"] = func(n_ Any) Any {
		n, _ := n_.(int)
		return strconv.Itoa(n)
	}
}
