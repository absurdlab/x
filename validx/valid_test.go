package validx_test

import (
	"errors"
	"github.com/absurdlab/x/validx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErr(t *testing.T) {
	var err = errors.New("")
	assert.Equal(t, err, validx.Err(err)())
}

func TestAnyErr(t *testing.T) {
	var err = errors.New("err")
	assert.Equal(t, err, validx.AnyErr(validx.Err(nil), validx.Err(err)))
}

func TestValidStr(t *testing.T) {
	assert.Nil(t, validx.ValidStr("any", func(s string) error {
		return nil
	})())
	assert.NotNil(t, validx.ValidStr("any", func(s string) error {
		return errors.New("err")
	})())
}

func TestNonEmptyStr(t *testing.T) {
	var err = errors.New("err")
	assert.Equal(t, err, validx.NonEmptyStr("", validx.Err(err))())
	assert.Nil(t, validx.NonEmptyStr("foo", validx.Err(err))())
}

func TestNonEmptyStrArr(t *testing.T) {
	var err = errors.New("err")
	assert.Equal(t, err, validx.NonEmptyStrArr([]string{}, validx.Err(err))())
	assert.Equal(t, err, validx.NonEmptyStrArr(nil, validx.Err(err))())
	assert.Nil(t, validx.NonEmptyStrArr([]string{"foo"}, validx.Err(err))())
}
