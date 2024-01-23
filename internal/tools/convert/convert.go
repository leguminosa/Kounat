package convert

import "strconv"

func ToInt(v interface{}) int {
	switch v := v.(type) {
	case int:
		return v
	case int8:
		return int(v)
	case int16:
		return int(v)
	case int32:
		return int(v)
	case int64:
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	case string:
		result, _ := strconv.Atoi(v)
		return result
	case []byte:
		return ToInt(string(v))
	case bool:
		if v {
			return 1
		}
		return 0
	default:
		return 0
	}
}
