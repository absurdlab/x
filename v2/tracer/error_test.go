package tracer_test

import (
	"errors"
	"github.com/absurdlab/x/v2/tracer"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	e1 = errors.New("e1")
	e2 = errors.New("e2")
	e3 = errors.New("e3")
)

var (
	traceSite1 = func() error {
		return tracer.Touch(e1)
	}
	traceSite2 = func(err error) error {
		return tracer.Touch(err)
	}
	traceSite3 = func(err error) error {
		return tracer.Touch(e2).Wrap(err)
	}
)

func TestTrace(t *testing.T) {
	var err error
	{
		err = traceSite1()
		err = traceSite2(err)
		err = traceSite3(err)
	}

	traces := tracer.Normalize(err).Traces()

	assert.Len(t, traces, 3)

	assert.Equal(t, "/lang/tracer/error_test.go", traces[0].File)
	assert.Equal(t, "/lang/tracer/error_test.go", traces[1].File)
	assert.Equal(t, "/lang/tracer/error_test.go", traces[2].File)

	assert.Greater(t, traces[0].Line, traces[1].Line)
	assert.Greater(t, traces[1].Line, traces[2].Line)
}

func TestUnwrap(t *testing.T) {
	var err error
	{
		err = traceSite1()
		err = traceSite2(err)
		err = traceSite3(err)
	}

	assert.True(t, errors.Is(err, e1))
	assert.True(t, errors.Is(err, e2))
	assert.False(t, errors.Is(err, e3))
}
