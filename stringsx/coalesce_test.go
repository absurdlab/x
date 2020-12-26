package stringsx_test

import (
	"github.com/absurdlab/x/stringsx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCoalesce(t *testing.T) {
	cases := []struct {
		name     string
		elements []string
		result   string
	}{
		{name: "single string", elements: []string{"foo"}, result: "foo"},
		{name: "multiple string", elements: []string{"", "foo", "bar"}, result: "foo"},
		{name: "no string", elements: []string{}, result: ""},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.result, stringsx.Coalesce(c.elements...))
		})
	}
}
