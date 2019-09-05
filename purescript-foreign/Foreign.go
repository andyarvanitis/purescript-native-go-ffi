package purescript_foreign

import (
	. "github.com/purescript-native/go-runtime"
)

func init() {
	exports := Foreign("Foreign")

	exports["unsafeToForeign"] = func(value Any) Any {
		return value
	}

	exports["unsafeFromForeign"] = func(value Any) Any {
		return value
	}

	exports["typeOf"] = func(value Any) Any {
		switch value.(type) {
		case string:
			return "string"

		case float64:
			return "number"

		case int:
			return "number"

		case bool:
			return "boolean"

		case Dict:
			return "object"

		case Any:
			return "object"

		case nil:
			return "object"

		default:
			return "undefined"
		}
	}

	exports["tagOf"] = func(value Any) Any {
		switch value.(type) {
		case string:
			return "String"

		case float64:
			return "Number"

		case int:
			return "Number"

		case bool:
			return "Boolean"

		case Dict:
			return "Object"

		case Any:
			return "Array"

		case nil:
			return "Null"

		default:
			return "Undefined"
		}
	}

	exports["isNull"] = func(value Any) Any {
		return value == nil
	}

	exports["isUndefined"] = func(value Any) Any {
		return false
	}

	exports["isArray"] = func(value Any) Any {
		_, ok := value.([]Any)
		return ok
	}
}
