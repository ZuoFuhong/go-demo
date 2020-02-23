package testify

import (
	"container/list"
	"github.com/stretchr/testify/assert"
	"testing"
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
