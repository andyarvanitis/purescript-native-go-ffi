package purescript_prelude

import (
	. "github.com/purescript-native/go-runtime"
)

func init() {
	exports := Foreign("Record.Unsafe")

	exports["unsafeSet"] = func(label_ Any) Any {
		return func(value Any) Any {
			return func(rec_ Any) Any {
				label, _ := label_.(string)
				rec, _ := rec_.(Dict)
				copy := make(Dict)
				for key, val := range rec {
					copy[key] = val
				}
				copy[label] = value
				return copy
			}
		}
	}

}
