package purescript_foreign

import (
	. "github.com/purescript-native/go-runtime"
)

func init() {
	exports := Foreign("Foreign.Index")

	exports["unsafeReadPropImpl"] = func(f, s, key_, value_ Any) Any {
		value, ok := value_.(Dict)
		if !ok {
			return f
		}
		key := key_.(string)
		return Apply(s, value[key])
	}
}
