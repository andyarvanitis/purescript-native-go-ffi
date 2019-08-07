package purescript_random

import (
	"math/rand"
	. "purescript"
	"time"
)

func init() {
	exports := Foreign("Effect.Random")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	exports["random"] = func() Any {
		return r.Float64()
	}
}
