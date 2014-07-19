package tiled

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointToIndex(t *testing.T) {
	boundaries := NewRect(0, 0, 10, 10)

	assert.Equal(t, 0, PointToIndex(NewPoint(0, 0), boundaries))
	assert.Equal(t, 99, PointToIndex(NewPoint(9, 9), boundaries))
	assert.Equal(t, 37, PointToIndex(NewPoint(7, 3), boundaries))
}
