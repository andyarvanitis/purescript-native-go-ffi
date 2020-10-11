package purescript_global

import (
	"encoding/json"

	. "github.com/purescript-native/go-runtime"
)

func init() {
	exports := Foreign("Simple.JSON")

	exports["_parseJSON"] = func(text Any) Any {
		var result Any
		if err := json.Unmarshal([]byte(text.(string)), &result); err != nil {
			panic(err.Error())
		}
		return result
	}

	exports["_undefined"] = nil
}
