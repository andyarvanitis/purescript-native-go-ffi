package purescript_refs

import . "github.com/purescript-native/go-runtime"

func init() {
	exports := Foreign("Effect.Ref")

	exports["new"] = func(val Any) Any {
		return func() Any {
			ptr := new(Any)
			*ptr = val
			return ptr
		}
	}

	exports["read"] = func(ref_ Any) Any {
		return func() Any {
			ref, _ := ref_.(*Any)
			return *ref
		}
	}

	exports["modify'"] = func(f Any) Any {
		return func(ref_ Any) Any {
			return func() Any {
				ref, _ := ref_.(*Any)
				t, _ := Apply(f, *ref).(Dict)
				*ref = t["state"]
				return t["value"]
			}
		}
	}

	exports["write"] = func(a Any) Any {
		return func(ref_ Any) Any {
			return func() Any {
				ref, _ := ref_.(*Any)
				*ref = a
				return nil
			}
		}
	}

}
