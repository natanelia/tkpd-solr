package solr

import (
	"strconv"
)

type mapOfStrings map[string]interface{}

func NewMapOfStrings(m map[string]interface{}) mapOfStrings {
	mos := make(mapOfStrings)
	for k, _ := range m {
		mos[k] = m[k]
	}
	return mos
}

func (m mapOfStrings) GetString(params ...string) string {
	if len(params) == 0 {
		return ""
	} else {
		switch vv := m[params[0]].(type) {
		case string:
			//fmt.Println("STRING")
			return vv
		case map[string]interface{}:
			//fmt.Println("MAP")
			a := NewMapOfStrings(vv)
			return a.GetString(params[1:]...)
		case int:
			//fmt.Println("INT")
			return strconv.Itoa(vv)
		case []interface{}:
			a, _ := strconv.Atoi(params[1])
			switch xx := vv[a].(type) {
			case string:
				return xx
			case map[string]interface{}:
				b := NewMapOfStrings(xx)
				return b.GetString(params[2:]...)
			case int:
				return strconv.Itoa(xx)

			}
		default:
			return ""
		}
	}
	return ""
}

func (m mapOfStrings) GetMapToString(params ...string) map[string]string {
	if len(params) == 0 {
		result := make(map[string]string)
		for k, v := range m {
			switch vv := v.(type) {
			case string:
				result[k] = vv
			case int:
				result[k] = strconv.Itoa(vv)
			default:
				result[k] = ""
			}
		}
		return result
	} else if len(params) == 1 {
		switch mm := m[params[0]].(type) {
		case map[string]interface{}:
			result := make(map[string]string)
			for k, v := range mm {
				switch vv := v.(type) {
				case string:
					result[k] = vv
				case int:
					result[k] = strconv.Itoa(vv)
				default:
					result[k] = ""
				}
			}
			return result
		}
	} else {
		switch vv := m[params[0]].(type) {
		case map[string]interface{}:
			a := NewMapOfStrings(vv)
			return a.GetMapToString(params[1:]...)
		case []interface{}:
			a, _ := strconv.Atoi(params[1])
			switch xx := vv[a].(type) {
			case map[string]interface{}:
				b := NewMapOfStrings(xx)
				if len(params) > 1 {
					return b.GetMapToString(params[2:]...)
				} else {
					return make(map[string]string)
				}
			default:
				return make(map[string]string)
			}
		default:
			return make(map[string]string)
		}
	}

	return make(map[string]string)
}

func (m mapOfStrings) GetMapToInterface(params ...string) map[string]interface{} {
	if len(params) == 0 {
		result := make(map[string]interface{})
		for k, v := range m {
			result[k] = v
		}
		return result
	} else if len(params) == 1 {
		switch mm := m[params[0]].(type) {
		case map[string]interface{}:
			result := make(map[string]interface{})
			for k, v := range mm {
				result[k] = v
			}
			return result
		default:
			return make(map[string]interface{})
		}
	} else {
		switch vv := m[params[0]].(type) {
		case map[string]interface{}:
			a := NewMapOfStrings(vv)
			return a.GetMapToInterface(params[1:]...)
		case []interface{}:
			a, _ := strconv.Atoi(params[1])
			switch xx := vv[a].(type) {
			case map[string]interface{}:
				b := NewMapOfStrings(xx)
				if len(params) > 1 {
					return b.GetMapToInterface(params[2:]...)
				} else {
					return make(map[string]interface{})
				}
			default:
				return make(map[string]interface{})
			}
		default:
			return make(map[string]interface{})
		}
	}

	return make(map[string]interface{})
}

func (m mapOfStrings) GetArrayOfInterface(params ...string) []interface{} {
	if len(params) == 0 {
		result := make([]interface{}, len(m))
		i := 0
		for _, v := range m {
			result[i] = v
			i++
		}
		return result
	} else if len(params) == 1 {
		switch mm := m[params[0]].(type) {
		case []interface{}:
			result := make([]interface{}, len(mm))
			for i, v := range mm {
				result[i] = v
			}
			return result
		}
	} else {
		switch vv := m[params[0]].(type) {
		case map[string]interface{}:
			a := NewMapOfStrings(vv)
			return a.GetArrayOfInterface(params[1:]...)
		case []interface{}:
			a, _ := strconv.Atoi(params[1])
			switch xx := vv[a].(type) {
			case map[string]interface{}:
				b := NewMapOfStrings(xx)
				if len(params) > 1 {
					return b.GetArrayOfInterface(params[2:]...)
				} else {
					return make([]interface{}, 0)
				}
			default:
				return make([]interface{}, 0)
			}
		default:
			return make([]interface{}, 0)
		}
	}

	return make([]interface{}, 0)
}
