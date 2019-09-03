package purescript_record

import (
	. "github.com/purescript-native/go-runtime"
)

func init() {
	exports := Foreign("Record.Builder")

	exports["copyRecord"] = func(rec_ Any) Any {
		rec := rec_.(map[string]Any)
		recCopy := make(map[string]Any)
		for key, value := range rec {
			recCopy[key] = value
		}
		return recCopy
	}

	exports["unsafeInsert"] = func(l_ Any) Any {
		return func(a_ Any) Any {
			return func(rec_ Any) Any {
				l := l_.(string)
				rec := rec_.(map[string]Any)
				rec[l] = a_
				return rec
			}
		}
	}
}
