package tiled

import (
	"fmt"
)

// Provides a view on a map that consists of well defined boundaries provided via
// the X(), Y(), Width(), and Height() methods, as well as access to the Layers
// within these boundaries.
// Map coordinates are always zero based.
type Map interface {
	Rect              // Embeds information about the position and size of the Map
	Layers() []Layer  // Return all layers
	Sub(Rect) Map     // Returns a subset of the map described by the given Rect
	Tileset() Tileset // Returns the tileset associated with this map.
}

type tmxMap struct {
	Rect
	layers  []Layer
	tileset Tileset
}

func (m *tmxMap) Layers() []Layer {
	return m.layers
}

func (m *tmxMap) Sub(r Rect) Map {
	layers := make([]Layer, len(m.layers))

	for i := range m.layers {
		layers[i] = m.layers[i].Sub(r)
	}

	return &tmxMap{
		Rect:   r,
		layers: layers,
	}
}

func (m *tmxMap) Tileset() Tileset {
	return m.tileset
}

func (m *tmxMap) String() string {
	return fmt.Sprintf("tmxMap{width: %d, height: %d}", m.Width(), m.Height())
}
