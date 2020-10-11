package purescript_strings

import (
	. "github.com/purescript-native/go-runtime"
	"unicode/utf8"
)

func init() {
	exports := Foreign("Data.String.Unsafe")

	exports["charAt"] = func(i_ Any) Any {
		return func(s_ Any) Any {
			i, _ := i_.(int)
			s, _ := s_.(string)
			runes := []rune(s)
			if i < 0 || i >= len(runes) {
				panic("Data.String.Unsafe.charAt: Invalid index.")
			}
			return string(runes[i : i+1])
		}
	}

	exports["char"] = func(s Any) Any {
		if utf8.RuneCountInString(s.(string)) != 1 {
			panic("Data.String.Unsafe.char: Expected string of length 1.")
		}
		return s
	}

}
