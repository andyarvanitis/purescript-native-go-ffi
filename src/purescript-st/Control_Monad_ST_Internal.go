package purescript_st

import . "purescript"

func init() {
	exports := Foreign("Control.Monad.ST.Internal")

	exports["map_"] = func(f Any) Any {
		return func(a Any) Any {
			return func() Any {
				return Apply(f, Run(a))
			}
		}
	}

	exports["run"] = func(f Any) Any {
		return Run(f)
	}

	exports["while"] = func(f Any) Any {
		return func(a Any) Any {
			return func() Any {
				for Run(f).(bool) {
					Run(a)
				}
				return nil
			}
		}
	}

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

	exports["write"] = func(a Any) Any {
		return func(ref Any) Any {
			return func() Any {
				ref.(Map)["value"] = a
				return a
			}
		}
	}
}
