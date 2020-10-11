package purescript_strings

import (
	// "fmt"
	. "github.com/purescript-native/go-runtime"
	"strings"
	"unicode/utf8"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func init() {
	exports := Foreign("Data.String.CodeUnits")

	exports["fromCharArray"] = func(a Any) Any {
		var buf strings.Builder
		for _, c_ := range a.([]Any) {
			c, _ := c_.(string)
			buf.WriteString(c)
		}
		return buf.String()
	}

	exports["toCharArray"] = func(s Any) Any {
		strs := strings.Split(s.(string), "")
		result := make([]Any, 0, len(strs))
		for _, str := range strs {
			result = append(result, str)
		}
		return result
	}

	exports["singleton"] = func(s Any) Any {
		return s
	}

	exports["_charAt"] = func(just Any) Any {
		return func(nothing Any) Any {
			return func(i_ Any) Any {
				return func(s_ Any) Any {
					i, _ := i_.(int)
					s, _ := s_.(string)
					runes := []rune(s)
					if i >= 0 && i < len(runes) {
						return Apply(just, string(runes[i]))
					}
					return nothing
				}
			}
		}
	}

	exports["_toChar"] = func(just Any) Any {
		return func(nothing Any) Any {
			return func(s Any) Any {
				if utf8.RuneCountInString(s.(string)) == 1 {
					return Apply(just, s)
				}
				return nothing
			}
		}
	}

	exports["length"] = func(s Any) Any {
		return utf8.RuneCountInString(s.(string))
	}

	exports["countPrefix"] = func(p Any) Any {
		return func(s Any) Any {
			runes := []rune(s.(string))
			i := 0
			for _, c := range runes {
				b, _ := Apply(p, string(c)).(bool)
				if !b {
					break
				}
				i++
			}
			return i
		}
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
						return Apply(just, utf8.RuneCountInString(s[:i]))
					}
				}
			}
		}
	}

	exports["_indexOfStartingAt"] = func(just Any) Any {
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
						substr := string(runes[startAt:])
						i := strings.Index(substr, x)
						if i == -1 {
							return nothing
						} else {
							return Apply(just, utf8.RuneCountInString(substr[:i])+startAt)
						}
						return nothing
					}
				}
			}
		}
	}

	exports["_lastIndexOf"] = func(just Any) Any {
		return func(nothing Any) Any {
			return func(x_ Any) Any {
				return func(s_ Any) Any {
					x, _ := x_.(string)
					s, _ := s_.(string)
					i := strings.LastIndex(s, x)
					if i == -1 {
						return nothing
					} else {
						return Apply(just, utf8.RuneCountInString(s[:i]))
					}
				}
			}
		}
	}

	exports["_lastIndexOfStartingAt"] = func(just Any) Any {
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
						substr := string(runes[:min(startAt+utf8.RuneCountInString(x), len(runes))])
						i := strings.LastIndex(substr, x)
						if i == -1 {
							return nothing
						} else {
							return Apply(just, utf8.RuneCountInString(substr[:i]))
						}
						return nothing
					}
				}
			}
		}
	}

	exports["take"] = func(n_ Any) Any {
		return func(s_ Any) Any {
			n, _ := n_.(int)
			if n < 0 {
				return ""
			}
			s, _ := s_.(string)
			runes := []rune(s)
			return string(runes[:min(n, len(runes))])
		}
	}

	exports["drop"] = func(n_ Any) Any {
		return func(s_ Any) Any {
			n, _ := n_.(int)
			s, _ := s_.(string)
			if n < 0 {
				return s
			}
			runes := []rune(s)
			return string(runes[min(n, len(runes)):])
		}
	}

	exports["_slice"] = func(b_ Any) Any {
		return func(e_ Any) Any {
			return func(s_ Any) Any {
				b, _ := b_.(int)
				e, _ := e_.(int)
				s, _ := s_.(string)
				runes := []rune(s)
				if b < 0 {
					b = max(len(runes)+b, 0)
				}
				if e < 0 {
					e = max(len(runes)+e, 0)
				}
				return string(runes[b:e])
			}
		}
	}

	exports["splitAt"] = func(n_ Any) Any {
		return func(s_ Any) Any {
			n, _ := n_.(int)
			s, _ := s_.(string)
			if n < 0 {
				return Dict{"before": "", "after": s}
			} else {
				runes := []rune(s)
				before := runes[:min(n, len(runes))]
				after := runes[min(n, len(runes)):]
				return Dict{"before": string(before), "after": string(after)}
			}
		}
	}

}
