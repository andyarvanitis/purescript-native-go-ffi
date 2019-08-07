package purescript_prelude

import . "github.com/purescript-native/go-runtime"

func ordImpl(flt func(Any, Any) bool) Fn {
	return func(lt Any) Any {
		return func(eq Any) Any {
			return func(gt Any) Any {
				return func(x Any) Any {
					return func(y Any) Any {
						if flt(x, y) {
							return lt
						} else if x == y {
							return eq
						} else {
							return gt
						}
					}
				}
			}
		}
	}
}

func init() {
	exports := Foreign("Data.Ord")

	exports["ordBooleanImpl"] = ordImpl(func(a Any, b Any) bool {
		ai := 0
		if a.(bool) {
			ai = 1
		}
		bi := 0
		if b.(bool) {
			bi = 1
		}
		return ai < bi
	})

	exports["ordIntImpl"] = ordImpl(func(a Any, b Any) bool {
		return a.(int) < b.(int)
	})

	exports["ordNumberImpl"] = ordImpl(func(a Any, b Any) bool {
		return a.(float64) < b.(float64)
	})

	exports["ordStringImpl"] = ordImpl(func(a Any, b Any) bool {
		return a.(string) < b.(string)
	})

	exports["ordCharImpl"] = ordImpl(func(a Any, b Any) bool {
		return a.(string) < b.(string)
	})

}
