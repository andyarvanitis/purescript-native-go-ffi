package purescript_global

import (
	"encoding/json"

	. "github.com/purescript-native/go-runtime"
)

func init() {
	exports := Foreign("Global.Unsafe")

	exports["unsafeStringify"] = func(value Any) Any {
		blob, err := json.Marshal(value)
		if err != nil {
			panic(err.Error())
		}
		return string(blob)
	}
}
