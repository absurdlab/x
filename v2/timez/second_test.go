package timez_test

import (
	"github.com/absurdlab/x/v2/timez"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSeconds(t *testing.T) {
	var tenSeconds timez.Seconds = 10
	assert.Equal(t, 10*time.Second, tenSeconds.Duration())
}
