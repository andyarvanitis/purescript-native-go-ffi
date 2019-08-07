package purescript_prelude

import (
	. "github.com/purescript-native/go-runtime"
	"math"
	"unicode/utf8"
)

const maxInt = int((^uint(0)) >> 1)
const minInt = -maxInt - 1

func init() {
	exports := Foreign("Data.Bounded")

	exports["topInt"] = maxInt
	exports["bottomInt"] = minInt

	exports["topChar"] = string(utf8.MaxRune)
	exports["bottomChar"] = string('\u0000')

	exports["topNumber"] = math.MaxFloat64
	exports["bottomNumber"] = math.SmallestNonzeroFloat64
}
