package purescript_prelude

import . "github.com/purescript-native/go-runtime"

func init() {
	exports := Foreign("Data.Unit")

	exports["unit"] = nil
}
