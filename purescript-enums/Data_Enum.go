package purescript_enums

import (
	. "github.com/purescript-native/go-runtime"
	"unicode/utf8"
)

func init() {
	exports := Foreign("Data.Enum")

	exports["toCharCode"] = func(c Any) Any {
		r, size := utf8.DecodeRuneInString(c.(string))
		if size == 0 {
			return 0
		}
		return int(r)
	}

	exports["fromCharCode"] = func(i Any) Any {
		return string(rune(i.(int)))
	}

}
