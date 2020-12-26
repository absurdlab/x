package strconvx_test

import (
	"github.com/absurdlab/x/strconvx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseInt64OrDefault(t *testing.T) {
	cases := []struct {
		name    string
		content string
		result  int64
	}{
		{name: "number", content: "3600", result: 3600},
		{name: "not number", content: "foo", result: 0},
		{name: "empty", content: "", result: 0},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.result, strconvx.ParseInt64OrDefault(c.content, 0))
		})
	}
}
