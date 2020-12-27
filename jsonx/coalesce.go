package jsonx

import "encoding/json"

// Coalesce returns the first non-empty json.RawMessage, or an empty json.RawMessage if all of them is empty.
func Coalesce(raws ...json.RawMessage) json.RawMessage {
	for _, each := range raws {
		if !IsEmptyObject(each) {
			return each
		}
	}
	return json.RawMessage{'{', '}'}
}
