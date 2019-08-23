package purescript_strings

import (
	. "github.com/purescript-native/go-runtime"
	"strings"
)

func init() {
	exports := Foreign("Data.String.Common")

	exports["_localeCompare"] = func(lt Any) Any {
		return func(eq Any) Any {
			return func(gt Any) Any {
				return func(s1 Any) Any {
					return func(s2 Any) Any {
						var result = strings.Compare(s1.(string), s2.(string))
						if result < 0 {
							return lt
						} else if result > 0 {
							return gt
						} else {
							return eq
						}
					}
				}
			}
		}
	}

	exports["replace"] = func(s1 Any) Any {
		return func(s2 Any) Any {
			return func(s3 Any) Any {
				return strings.Replace(s3.(string), s1.(string), s2.(string), 1)
			}
		}
	}

	exports["replaceAll"] = func(s1 Any) Any {
		return func(s2 Any) Any {
			return func(s3 Any) Any {
				return strings.ReplaceAll(s3.(string), s1.(string), s2.(string))
			}
		}
	}

	exports["split"] = func(sep Any) Any {
		return func(s Any) Any {
			xs := strings.Split(s.(string), sep.(string))
			result := make([]Any, 0, len(xs))
			for _, x := range xs {
				result = append(result, x)
			}
			return result
		}
	}

	exports["toLower"] = func(s Any) Any {
		return strings.ToLower(s.(string))
	}

	exports["toUpper"] = func(s Any) Any {
		return strings.ToUpper(s.(string))
	}

	exports["trim"] = func(s Any) Any {
		return strings.TrimSpace(s.(string))
	}

	exports["joinWith"] = func(s Any) Any {
		return func(xs_ Any) Any {
			xs := xs_.([]Any)
			ss := make([]string, 0, len(xs))
			for _, x := range xs {
				ss = append(ss, x.(string))
			}
			return strings.Join(ss, s.(string))
		}
	}
}
