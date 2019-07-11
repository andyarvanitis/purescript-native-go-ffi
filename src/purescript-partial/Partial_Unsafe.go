package purescript_partial

import (
	"Partial_Unsafe"
	. "purescript"
)

func init() {
	exports := Partial_Unsafe.Foreign

	exports["unsafePartial"] = func(f Any) Any {
		return Run(f)
	}

}
