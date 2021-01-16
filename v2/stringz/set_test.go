package stringz_test

import (
	"encoding/json"
	"github.com/absurdlab/x/v2/stringz"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStableSet_Add(t *testing.T) {
	set := stringz.NewSet().Add("a", "b", "c")

	assert.Equal(t, 3, set.Len())

	assert.True(t, set.Contains("a"))
	assert.True(t, set.Contains("b"))
	assert.True(t, set.Contains("c"))

	assert.Equal(t, []string{"a", "b", "c"}, set.Array())
}

func TestStableSet_Remove(t *testing.T) {
	set := stringz.NewSet().Add("a", "b", "c").Remove("b")

	assert.Equal(t, 2, set.Len())

	assert.True(t, set.Contains("a"))
	assert.False(t, set.Contains("b"))
	assert.True(t, set.Contains("c"))

	assert.Equal(t, []string{"a", "c"}, set.Array())
}

func TestStableSet_JSON(t *testing.T) {
	set := stringz.NewSet().Add("a", "b", "c")

	raw, _ := set.MarshalJSON()
	assert.JSONEq(t, `["a","b","c"]`, string(raw))

	set2 := stringz.NewSet()
	assert.NoError(t, json.Unmarshal(raw, &set2))

	assert.True(t, set2.Contains("a"))
	assert.True(t, set2.Contains("b"))
	assert.True(t, set2.Contains("c"))
}

func TestStableSet_JSON2(t *testing.T) {
	type obj struct {
		Set *stringz.Set `json:"set"`
	}

	var v obj
	err := json.Unmarshal([]byte(`{"set": ["a", "b", "c"]}`), &v)
	assert.NoError(t, err)

	assert.True(t, v.Set.Contains("a"))
	assert.True(t, v.Set.Contains("b"))
	assert.True(t, v.Set.Contains("c"))
}
