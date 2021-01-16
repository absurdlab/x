package stringz

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiSet_MarshalJSON(t *testing.T) {
	type obj struct {
		ResponseTypes MultiSet `json:"response_types"`
	}

	raw, err := json.Marshal(obj{
		ResponseTypes: []*Set{
			NewSet().Add("code"),
			NewSet().Add("code id_token"),
		},
	})
	assert.NoError(t, err)
	assert.JSONEq(t, `{"response_types":["code","code id_token"]}`, string(raw))
}

func TestMultiSet_UnmarshalJSON(t *testing.T) {
	type obj struct {
		ResponseTypes MultiSet `json:"response_types"`
	}

	var v obj
	err := json.Unmarshal([]byte(`{"response_types":["code","code id_token"]}`), &v)
	assert.NoError(t, err)

	assert.Len(t, v.ResponseTypes, 2)
	assert.True(t, v.ResponseTypes[0].Equals(NewSet().Add("code")))
	assert.True(t, v.ResponseTypes[1].Equals(NewSet().Add("code", "id_token")))
}
