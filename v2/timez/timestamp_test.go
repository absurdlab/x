package timez_test

import (
	"github.com/absurdlab/x/v2/timez"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTimestamp(t *testing.T) {
	now := time.Now().Round(time.Second)
	ts := timez.ToTimestamp(now)
	assert.True(t, now.Equal(ts.Time()))
}
