package purescript_lazy

import (
	. "github.com/purescript-native/go-runtime"
	"sync"
)

func init() {
	exports := Foreign("Data.Lazy")

	exports["defer"] = func(thunk Any) Any {
		var once sync.Once
		var value Any = nil
		return func() Any {
			once.Do(func() {
				value = Apply(thunk, nil)
			})
			return value
		}
	}

	exports["force"] = func(l Any) Any {
		return Run(l)
	}
}
