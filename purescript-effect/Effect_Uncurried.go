package purescript_effect

import (
	. "github.com/purescript-native/go-runtime"
)

func init() {
	exports := Foreign("Effect.Uncurried")

	exports["runEffectFn1"] = func(fn_ Any) Any {
		return func(a_ Any) Any {
			return func() Any {
				return Apply(fn_, a_)
			}
		}
	}
}
