package purescript_global

import (
	. "github.com/purescript-native/go-runtime"
	"math"
)

func init() {
	exports := Foreign("Global")

	exports["infinity"] = math.Inf(1)

}
