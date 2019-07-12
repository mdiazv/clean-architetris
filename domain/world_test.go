package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWorld(t *testing.T) {
	tw := MakeTetrisWorld()
	w, h := tw.Dimensions()
	assert.Equal(t, 10, w)
	assert.Equal(t, 20, h)
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			assert.True(t, tw.IsFree(x, y))
		}
	}
}
