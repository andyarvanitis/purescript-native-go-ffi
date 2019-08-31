package purescript_strings

import (
	. "github.com/purescript-native/go-runtime"
	"strings"
	"unicode/utf8"
)

func init() {
	exports := Foreign("Data.String.CodeUnits")

	exports["length"] = func(s_ Any) Any {
		s, _ := s_.(string)
		return utf8.RuneCountInString(s)
	}

	exports["_indexOf"] = func(just Any) Any {
		return func(nothing Any) Any {
			return func(x_ Any) Any {
				return func(s_ Any) Any {
					x, _ := x_.(string)
					s, _ := s_.(string)
					i := strings.Index(s, x)
					if i == -1 {
						return nothing
					} else {
						return Apply(just, i)
					}
				}
			}
		}
	}

	exports["_indexOf'"] = func(just Any) Any {
		return func(nothing Any) Any {
			return func(x_ Any) Any {
				return func(startAt_ Any) Any {
					return func(s_ Any) Any {
						x, _ := x_.(string)
						startAt, _ := startAt_.(int)
						s, _ := s_.(string)
						runes := []rune(s)
						if startAt < 0 || startAt > len(runes) {
							return nothing
						}
						r, _ := utf8.DecodeRuneInString(x)
						for i, c := range runes[startAt:] {
							if c == r {
								return Apply(just, i)
							}
						}
						return nothing
					}
				}
			}
		}
	}

	exports["drop"] = func(n_ Any) Any {
		return func(s_ Any) Any {
			n, _ := n_.(int)
			s, _ := s_.(string)
			return s[n:]
		}
	}
}
