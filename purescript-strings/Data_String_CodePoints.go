package purescript_strings

import (
	. "github.com/purescript-native/go-runtime"
	"unicode/utf8"
)

func init() {
	exports := Foreign("Data.String.CodePoints")

	exports["_unsafeCodePointAt0"] = func(_ Any) Any {
		return func(str_ Any) Any {
			str, _ := str_.(string)
			r, _ := utf8.DecodeRuneInString(str)
			if r == utf8.RuneError {
				return Undefined
			}
			return int(r)
		}
	}

	exports["_codePointAt"] = func(_ Any) Any {
		return func(just Any) Any {
			return func(nothing Any) Any {
				return func(unsafeCodePointAt0 Any) Any {
					return func(index_ Any) Any {
						return func(s_ Any) Any {
							index, _ := index_.(int)
							s, _ := s_.(string)
							runes := []rune(s)
							if index < 0 || index >= len(runes) {
								return nothing
							}
							result := Apply(unsafeCodePointAt0, string(runes[index]))
							return Apply(just, result)
						}
					}
				}
			}
		}
	}

	exports["_countPrefix"] = func(_ Any) Any {
		return func(unsafeCodePointAt0 Any) Any {
			return func(pred Any) Any {
				return func(s_ Any) Any {
					s, _ := s_.(string)
					runes := []rune(s)
					i := 0
					for _, r := range runes {
						cp := Apply(unsafeCodePointAt0, string(r))
						cond, _ := Apply(pred, cp).(bool)
						if !cond {
							return i
						}
						i++
					}
					return i
				}
			}
		}
	}

	exports["_singleton"] = func(_ Any) Any {
		return func(c_ Any) Any {
			c, _ := c_.(int)
			return string(rune(c))
		}
	}

	exports["_take"] = func(_ Any) Any {
		return func(n_ Any) Any {
			return func(s_ Any) Any {
				n, _ := n_.(int)
				s, _ := s_.(string)
				if n < 0 {
					return ""
				}
				runes := []rune(s)
				// return string(runes[:n])
				return string(runes[:min(n, len(runes))])
			}
		}
	}

	exports["_toCodePointArray"] = func(_ Any) Any {
		return func(unsafeCodePointAt0 Any) Any {
			return func(s_ Any) Any {
				s, _ := s_.(string)
				runes := []rune(s)
				result := make([]Any, 0, len(runes))
				for _, r := range runes {
					result = append(result, Apply(unsafeCodePointAt0, string(r)))
				}
				return result
			}
		}
	}
}
