package stringsx_test

import (
	"github.com/absurdlab/x/stringsx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDelimited(t *testing.T) {
	cases := []struct {
		name    string
		content string
		expect  func(t *testing.T, fields []string)
	}{
		{
			name:    "empty string",
			content: "",
			expect: func(t *testing.T, fields []string) {
				assert.Len(t, fields, 0)
			},
		},
		{
			name:    "space delimited",
			content: "foo bar",
			expect: func(t *testing.T, fields []string) {
				assert.Len(t, fields, 2)
				assert.Contains(t, fields, "foo")
				assert.Contains(t, fields, "bar")
			},
		},
		{
			name:    "multiple space delimited",
			content: "foo  bar baz",
			expect: func(t *testing.T, fields []string) {
				assert.Len(t, fields, 3)
				assert.Contains(t, fields, "foo")
				assert.Contains(t, fields, "bar")
				assert.Contains(t, fields, "baz")
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			c.expect(t, stringsx.Delimited(c.content).BySpace())
		})
	}
}
