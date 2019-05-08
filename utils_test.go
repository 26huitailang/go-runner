package main

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	x1, y1 := 1, 1
	ret1 := 2
	x2, y2 := 2, 5
	ret2 := 7
	x3, y3 := 19, 20
	ret3 := 39
	assert.Equal(t, Add(x1, y1), ret1, "fail")
	assert.Equal(t, Add(x2, y2), ret2, "fail")
	assert.Equal(t, Add(x3, y3), ret3, "fail")
}
