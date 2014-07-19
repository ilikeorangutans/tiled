package tiled

import (
	"github.com/ilikeorangutans/tiled/tmx"
)

type Layer interface {
	Rect
	Name() string
	TileAt(Point) Tile
	Sub(Rect) Layer
}

// Layer implementation backed by a TMX layer. Layers can be "sliced" into smaller
// portions of itself using the Sub() method. A sliced layer uses the same slice of
// Tiles that backs its parent.
type tmxLayer struct {
	Rect          // Defines the boundaries of this layer. If this layer is a sub of another layer, the Rect's origin might not be 0/0.
	name   string // Name of the layer.
	tiles  []Tile // The actual tiles as an linear array.
	parent Layer
}

func (l *tmxLayer) Name() string {
	return l.name
}

// Returns the tile at the given point, using the upper left edge of the layer as
// the origin.
func (l *tmxLayer) TileAt(p Point) Tile {
	pointToLookup := p
	boundaries := l.Rect

	if l.parent != nil {
		boundaries = l.parent
		pointToLookup = NewPoint(l.X()+p.X(), l.Y()+p.Y())
	}

	return l.tiles[PointToIndex(pointToLookup, boundaries)]
}

func (l *tmxLayer) Sub(rect Rect) Layer {
	return &tmxLayer{
		Rect:   rect,
		tiles:  l.tiles,
		parent: l,
	}
}

// Creates a new Layer instance from the given TMX data
func LayerFromTMX(boundaries Rect, l tmx.Layer) Layer {

	tiles := make([]Tile, Area(boundaries))

	var x, y int
	for i := range l.Data.Tiles {
		x = i % boundaries.Width()
		tmxTile := l.Data.Tiles[i]
		tiles[i] = TileFromTMX(&tmxTile, NewPoint(x, y))

		endOfRow := x%boundaries.Width() == boundaries.Width()-1
		if endOfRow {
			y++
		}
	}

	return &tmxLayer{
		Rect:  boundaries,
		name:  l.Name,
		tiles: tiles,
	}
}
