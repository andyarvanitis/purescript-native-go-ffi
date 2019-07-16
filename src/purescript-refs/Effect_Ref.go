package purescript_refs

import . "purescript"

func init() {
	exports := Foreign("Effect.Ref")

	exports["new"] = func(val Any) Any {
		return func() Any {
			return Map{"value": val}
		}
	}

	exports["read"] = func(ref Any) Any {
		return func() Any {
			return ref.(Map)["value"]
		}
	}

	exports["modify'"] = func(f Any) Any {
		return func(ref_ Any) Any {
			return func() Any {
				ref, _ := ref_.(Map)
				t, _ := Apply(f, ref["value"]).(Map)
				ref["value"] = t["state"]
				return t["value"]
			}
		}
	}

	exports["write"] = func(val Any) Any {
		return func(ref Any) Any {
			return func() Any {
				ref.(Map)["value"] = val
				return nil
			}
		}
	}

}
