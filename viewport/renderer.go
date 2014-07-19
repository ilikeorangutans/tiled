package viewport

import (
	"github.com/ilikeorangutans/tiled"
	"github.com/veandco/go-sdl2/sdl"
)

func NewRenderer(m tiled.Map, viewport Viewport) *MapRenderer {

	// Calculate how many tiles we can display in the current viewport:
	// TODO: this should probably go into the viewport as the viewport might be
	// resized (and then we'd need to recalculate these values)
	widthInTiles := viewport.Width() >> 5
	if viewport.Width()%32 > 0 {
		widthInTiles++
	}
	heightInTiles := viewport.Height() >> 5
	if viewport.Height()%32 > 0 {
		heightInTiles++
	}

	return &MapRenderer{
		m:             m,
		viewport:      viewport,
		widthInTiles:  widthInTiles,
		heightInTiles: heightInTiles,
	}
}

type MapRenderer struct {
	viewport                    Viewport
	tileWidth, tileHeight       int
	m                           tiled.Map
	widthInTiles, heightInTiles int // number of tiles required to fill the viewport horizontally and vertically
}

// Render the map with the current viewport to the given surface.
func (r *MapRenderer) Render(surface *sdl.Surface) {

}
