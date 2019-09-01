package purescript_global

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

}

// exports.fromCharCode = function (c) {
//   return String.fromCharCode(c);
// };