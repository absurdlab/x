package stringsx_test

import (
	"github.com/absurdlab/x/stringsx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestSet(t *testing.T) {
	suite.Run(t, new(SetTestSuite))
}

type SetTestSuite struct {
	suite.Suite
}

func (s *SetTestSuite) TestLen() {
	assert.Equal(s.T(), 0, stringsx.NewSet().Len())
	assert.Equal(s.T(), 1, stringsx.NewSet("foo").Len())
	assert.Equal(s.T(), 2, stringsx.NewSet("foo", "bar", "foo").Len())
}

func (s *SetTestSuite) TestFirst() {
	assert.Equal(s.T(), "foo", stringsx.NewSet("foo").First())
}

func (s *SetTestSuite) TestContainsAll() {
	assert.True(s.T(), stringsx.NewSet("foo", "bar", "baz").ContainsAll(stringsx.NewSet("foo", "baz")))
	assert.False(s.T(), stringsx.NewSet("foo", "bar", "baz").ContainsAll(stringsx.NewSet("foo", "invalid")))
}

func (s *SetTestSuite) TestContainsAny() {
	assert.True(s.T(), stringsx.NewSet("foo", "bar", "baz").ContainsAny(stringsx.NewSet("foo")))
	assert.False(s.T(), stringsx.NewSet("foo", "bar", "baz").ContainsAny(stringsx.NewSet("invalid")))
}

func (s *SetTestSuite) TestEquals() {
	assert.True(s.T(), stringsx.NewSet("foo", "bar").Equals(stringsx.NewSet("foo", "bar")))
	assert.False(s.T(), stringsx.NewSet("foo", "bar").Equals(stringsx.NewSet("foo")))
}

func (s *SetTestSuite) TestAddAll() {
	set := stringsx.NewSet()
	assert.False(s.T(), set.Contains("foo"))

	set.Add("foo")
	assert.True(s.T(), set.Contains("foo"))
}
