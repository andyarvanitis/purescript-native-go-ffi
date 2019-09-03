package purescript_foreign

import (
	. "github.com/purescript-native/go-runtime"
)

func init() {
	exports := Foreign("Foreign")

	exports["unsafeToForeign"] = func(value_ Any) Any {
		return value_
	}

	exports["unsafeFromForeign"] = func(value_ Any) Any {
		return value_
	}

	exports["typeOf"] = func(value_ Any) Any {
		switch value_.(type) {
		//NOTE: do not use Any interface in type switches
		case string:
			return "string"

		case float64:
			return "number"

		case int:
			return "number"

		case bool:
			return "boolean"

		case map[string]interface{}:
			return "object"

		case []interface{}:
			return "object"

		default:
			return "undefined"
		}
	}

	exports["tagOf"] = func(value_ Any) Any {
		switch value_.(type) {
		//NOTE: do not use Any interface in type switches
		case string:
			return "String"

		case float64:
			return "Number"

		case int:
			return "Number"

		case bool:
			return "Boolean"

		case map[string]interface{}:
			return "Object"

		case []interface{}:
			return "Array"

		default:
			return "Object"
		}
	}

	exports["isNull"] = func(value_ Any) Any {
		return value_ == nil
	}

	exports["isUndefined"] = func(value_ Any) Any {
		return false
	}

	exports["isArray"] = func(value_ Any) Any {
		//NOTE: do not use Any interface here
		_, ok := value_.([]interface{})
		return ok
	}
}
