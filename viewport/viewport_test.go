package viewport

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestViewportDefaults(t *testing.T) {
	v := NewViewport(800, 600, 0, 0, 32, 32)

	assert.Equal(t, 800, v.Width())
	assert.Equal(t, 600, v.Height())
	assert.Equal(t, 1024, v.AreaWidth())
	assert.Equal(t, 1024, v.AreaHeight())
}

func TestMoveShouldNotMoveOutsideOfBoundaries(t *testing.T) {
	v := NewViewport(800, 600, 0, 0, 32, 32)

	v.MoveBy(-10, -10)
	assert.Equal(t, 0, v.X())
	assert.Equal(t, 0, v.Y())

	v.MoveBy(1024, 1024)
	assert.Equal(t, 192, v.X())
	assert.Equal(t, 392, v.Y())

}

func TestVisibleTiles(t *testing.T) {

	vp := NewViewport(100, 100, 0, 0, 32, 32)

	x1, y1, x2, y2 := vp.VisibleTiles()

	assert.Equal(t, 0, x1, "top left corner should be at origin")
	assert.Equal(t, 0, y1, "top left corner should be at origin")
	assert.Equal(t, 5, x2)
	assert.Equal(t, 5, y2)

	vp.MoveTo(5, 5)
	x1, y1, x2, y2 = vp.VisibleTiles()

	assert.Equal(t, 0, x1)
	assert.Equal(t, 0, y1)
	assert.Equal(t, 5, x2)
	assert.Equal(t, 5, y2)

	vp.MoveTo(32, 32)
	log.Printf("viewport is now at %d/%d", vp.X(), vp.Y())
	x1, y1, x2, y2 = vp.VisibleTiles()

	assert.Equal(t, 1, x1)
	assert.Equal(t, 1, y1)
	assert.Equal(t, 6, x2)
	assert.Equal(t, 6, y2)

	vp.MoveTo(1024, 1024)
	log.Printf("viewport is now at %d/%d", vp.X(), vp.Y())
	x1, y1, x2, y2 = vp.VisibleTiles()

	assert.Equal(t, 27, x1)
	assert.Equal(t, 27, y1)
	assert.Equal(t, 31, x2)
	assert.Equal(t, 31, y2)

}
