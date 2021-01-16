package stringz

import (
	"encoding/json"
	"github.com/absurdlab/x/v2/tracer"
	"strings"
)

// MultiSet represents a list of Set of composite element, where each element is a space delimited string. For
// example: ["code", "code token", "code id_token"]
type MultiSet []*Set

func (l *MultiSet) UnmarshalJSON(bytes []byte) error {
	rs := make([]string, 0)
	if err := json.Unmarshal(bytes, &rs); err != nil {
		return tracer.Touch(err)
	}

	var list []*Set
	for _, it := range rs {
		list = append(list, NewSet().Add(strings.Fields(it)...))
	}

	if list == nil {
		return nil
	}

	*l = list
	return nil
}

func (l MultiSet) MarshalJSON() ([]byte, error) {
	var rs []string
	for _, it := range l {
		rs = append(rs, strings.Join(it.Array(), " "))
	}

	if len(rs) == 0 {
		return json.Marshal([]string{})
	}

	raw, err := json.Marshal(rs)
	if err != nil {
		return nil, tracer.Touch(err)
	}

	return raw, nil
}
