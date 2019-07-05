package purescript_prelude

import (
	"Data_Bounded"
	"math"
	"unicode/utf8"
)

func init() {
	exports := Data_Bounded.Foreign

	exports["topInt"] = math.MaxInt32
	exports["bottomInt"] = math.MinInt32

	exports["topChar"] = string(utf8.MaxRune)
	exports["bottomChar"] = string('\u0000')

	exports["topNumber"] = math.MaxFloat64
	exports["bottomNumber"] = math.SmallestNonzeroFloat64
}
