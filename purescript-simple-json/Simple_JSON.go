package purescript_global

import (
	"encoding/json"

	. "github.com/purescript-native/go-runtime"
)

func init() {
	exports := Foreign("Simple.JSON")

	exports["_parseJSON"] = func(text_ Any) Any {
		text := text_.(string)
		var result Any
		err := json.Unmarshal([]byte(text), &result)
		if err != nil {
			panic(err.Error())
		}
		return result
	}
}
