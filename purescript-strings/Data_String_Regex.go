package purescript_strings

import (
	"fmt"
	. "github.com/purescript-native/go-runtime"
	"regexp"
	"strings"
	"unicode/utf8"
)

type regex_pair struct {
	regex  *regexp.Regexp
	global bool
}

func init() {
	exports := Foreign("Data.String.Regex")

	exports["regex'"] = func(left Any) Any {
		return func(right Any) Any {
			return func(s_ Any) Any {
				return func(flags_ Any) Any {
					s, _ := s_.(string)
					flags, _ := flags_.(string)
					global := false
					if strings.Contains(flags, "g") {
						global = true
						flags = strings.ReplaceAll(flags, "g", "")
						if flags != "" {
							flags = fmt.Sprintf("(?%s)", flags)
						}
					}
					r, err := regexp.Compile(flags + s)
					if err == nil {
						return Apply(right, regex_pair{r, global})
					} else {
						return Apply(left, err.Error())
					}
				}
			}
		}
	}

	exports["test"] = func(p_ Any) Any {
		return func(s_ Any) Any {
			p := p_.(regex_pair)
			r := p.regex
			s, _ := s_.(string)
			return r.MatchString(s)
		}
	}

	exports["_match"] = func(just Any) Any {
		return func(nothing Any) Any {
			return func(p_ Any) Any {
				return func(s_ Any) Any {
					p := p_.(regex_pair)
					r := p.regex
					s, _ := s_.(string)
					ms := r.FindAllString(s, -1)
					if ms == nil {
						return nothing
					}
					result := make([]Any, 0, len(ms))
					for _, m := range ms {
						if m == "" {
							result = append(result, nothing)
						} else {
							result = append(result, Apply(just, m))
						}
					}
					return Apply(just, result)
				}
			}
		}
	}

	exports["replace"] = func(p_ Any) Any {
		return func(s1_ Any) Any {
			return func(s2_ Any) Any {
				p := p_.(regex_pair)
				r := p.regex
				global := p.global

				s1, _ := s1_.(string)
				s2, _ := s2_.(string)

				if global {
					return r.ReplaceAllString(s2, s1)
				}

				found := r.FindString(s2)
				if found != "" {
					return strings.Replace(s2, found, s1, 1)
				}
				return s2
			}
		}
	}

	exports["replace'"] = func(p_ Any) Any {
		return func(f Any) Any {
			return func(s_ Any) Any {
				p := p_.(regex_pair)
				r := p.regex
				global := p.global
				s, _ := s_.(string)

				all := r.FindAllString(s, -1)
				submatches := make([]Any, 0, len(all))
				for _, submatch := range all {
					submatches = append(submatches, submatch)
				}

				frepl := func(str string) string {
					return Apply(f, str, submatches).(string)
				}

				if global {
					return r.ReplaceAllStringFunc(s, frepl)
				}

				found := r.FindString(s)
				if found != "" {
					fmt.Println(r.ReplaceAllStringFunc(s, frepl))
					return strings.Replace(s, found, frepl(r.String()), 1)
				}
				return s
			}
		}
	}

	exports["_search"] = func(just Any) Any {
		return func(nothing Any) Any {
			return func(p_ Any) Any {
				return func(s_ Any) Any {
					p := p_.(regex_pair)
					r := p.regex
					s, _ := s_.(string)
					found := r.FindStringIndex(s)
					if found == nil {
						return nothing
					}
					// TODO: is there a way to do this that is faster?
					return Apply(just, utf8.RuneCountInString(s[:found[0]]))
				}
			}
		}
	}

	exports["split"] = func(p_ Any) Any {
		return func(s_ Any) Any {
			p := p_.(regex_pair)
			r := p.regex
			s, _ := s_.(string)
			ss := r.Split(s, -1)
			result := make([]Any, 0, len(ss))
			for _, str := range ss {
				result = append(result, str)
			}
			return result
		}
	}
}
