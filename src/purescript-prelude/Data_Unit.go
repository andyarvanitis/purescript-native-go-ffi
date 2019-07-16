package purescript_prelude

import . "purescript"

func init() {
	exports := Foreign("Data.Unit")

	exports["unit"] = nil
}
