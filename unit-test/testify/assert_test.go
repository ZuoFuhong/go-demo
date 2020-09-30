package testify

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_assert(t *testing.T) {
	a := 2
	b := 3
	assert.Equal(t, a+b, 5, "They should be equal")
}

// require相比assert没有返回值，而是终止当前测试
func Test_require(t *testing.T) {
	var name = "dazuo"
	var age = 24
	require.Equal(t, "dazuo", name, "they should be equal")
	require.Equal(t, 24, age, "they should be equal")
}
