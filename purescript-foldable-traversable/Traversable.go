package purescript_foldable_traversable

import (
	. "github.com/purescript-native/go-runtime"
	"math"
)

func array1(a Any) Any {
	return []Any{a}
}

func array2(a Any) Any {
	return func(b Any) Any {
		return []Any{a, b}
	}
}

func array3(a Any) Any {
	return func(b Any) Any {
		return func(c Any) Any {
			return []Any{a, b, c}
		}
	}
}

func concat2(xs_ Any) Any {
	return func(ys_ Any) Any {
		xs, _ := xs_.([]Any)
		ys, _ := ys_.([]Any)
		return append(xs, ys...)
	}
}

func init() {
	exports := Foreign("Data.Traversable")

	exports["traverseArrayImpl"] = func(apply Any) Any {
		return func(_map Any) Any {
			return func(pure Any) Any {
				return func(f_ Any) Any {
					return func(array_ Any) Any {
						f, _ := f_.(Fn)
						array, _ := array_.([]Any)
						var _go func(int, int) Any
						_go = func(bot int, top int) Any {
							switch top - bot {
							case 0:
								return Apply(pure, []Any{})
							case 1:
								return Apply(_map, array1, f(array[bot]))
							case 2:
								return Apply(apply, Apply(_map, array2, f(array[bot])), f(array[bot+1]))
							case 3:
								return Apply(apply, Apply(apply, Apply(_map, array3, f(array[bot])), f(array[bot+1])), f(array[bot+2]))
							default:
								// This slightly tricky pivot selection aims to produce two
								// even-length partitions where possible.
								pivot := bot + int(math.Floor(float64(top-bot)/4))*2
								return Apply(apply, Apply(_map, concat2, _go(bot, pivot)), _go(pivot, top))
							}
						}
						return _go(0, len(array))
					}
				}
			}
		}
	}

}
