package purescript_effect

import (
	. "github.com/purescript-native/go-runtime"
)

func init() {
	exports := Foreign("Effect")

	exports["pureE"] = func(a Any) Any {
		return func() Any {
			return a
		}
	}

	exports["bindE"] = func(a Any) Any {
		return func(f Any) Any {
			return func() Any {
				return Run(Apply(f, Run(a)))
			}
		}
	}

	exports["untilE"] = func(f Any) Any {
		return func() Any {
			for !(Run(f).(bool)) {
			}
			return nil
		}
	}

	exports["whileE"] = func(f Any) Any {
		return func(a Any) Any {
			return func() Any {
				for Run(f).(bool) {
					Run(a)
				}
				return nil
			}
		}
	}

	exports["forE"] = func(lo Any) Any {
		return func(hi Any) Any {
			return func(f Any) Any {
				return func() Any {
					for i := lo.(int); i < hi.(int); i++ {
						Run(Apply(f, i))
					}
					return nil
				}
			}
		}
	}

	exports["foreachE"] = func(as Any) Any {
		return func(f Any) Any {
			return func() Any {
				for _, a := range as.([]Any) {
					Run(Apply(f, a))
				}
				return nil
			}
		}
	}

}
