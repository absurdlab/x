package jsonx

import "encoding/json"

func IsNull(raw json.RawMessage) bool {
	return len(raw) == 4 &&
		raw[0] == 'n' &&
		raw[1] == 'u' &&
		raw[2] == 'l' &&
		raw[3] == 'l'
}

func IsEmptyObject(raw json.RawMessage) bool {
	if len(raw) == 0 {
		return true
	}

	depth := 0
	for _, r := range raw {
		switch r {
		case '{':
			depth = depth + 1
			if depth > 1 {
				return false
			}
		case '}':
			depth = depth - 1
		case ' ':
			continue
		default:
			return false
		}
	}

	return false
}

// MustEncode encodes the given object and panics on error. Exceptions where given to
// type of []byte, json.RawMessage and string, as these types are returned as []byte
// without encoding.
func MustEncode(obj interface{}) []byte {
	switch obj.(type) {
	case []byte:
		return obj.([]byte)
	case json.RawMessage:
		return obj.([]byte)
	case string:
		return []byte(obj.(string))
	default:
		raw, err := json.Marshal(obj)
		if err != nil {
			panic(err)
		}
		return raw
	}
}
