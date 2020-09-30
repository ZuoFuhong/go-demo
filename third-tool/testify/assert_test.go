package testify

import (
	"container/list"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func Test_assert(t *testing.T) {
	assertor := assert.New(t)
	a := 2
	b := 3
	assertor.Equal(1, b-a, "they should be equal")

	var l = list.New()
	assertor.Nil(l.Front())

	assert.True(t, true, "True is true")
}

// require没有返回值，而是终止当前测试
func Test_require(t *testing.T) {
	var name = "dazuo"
	var age = 24
	require.Equal(t, "dazuo", name, "they should be equal")
	require.Equal(t, 24, age, "they should be equal")
}
