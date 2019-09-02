package purescript_enums

import (
	. "github.com/purescript-native/go-runtime"
	"unicode/utf8"
)

func init() {
	exports := Foreign("Data.Enum")

	exports["toCharCode"] = func(c_ Any) Any {
		c, _ := c_.(string)
		r, _ := utf8.DecodeRuneInString(c)
		return int(r)
	}

	exports["fromCharCode"] = func(i_ Any) Any {
		i, _ := i_.(int)
		return string(rune(i))
	}

}
