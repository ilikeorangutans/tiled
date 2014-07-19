package tiled

import (
	"github.com/ilikeorangutans/tiled/tmx"
)

// A single tile on a Map.
type Tile interface {
	Point
	Type() TileType
}

type TileType interface {
	Gid() int
}

func TileFromTMX(tmx *tmx.Tile, position Point) Tile {
	// TODO: build some kind of tile catalog?!
	return &tmxTile{
		Point: position,
	}
}

type tmxTile struct {
	Point
}

func (t *tmxTile) Type() TileType {
	return nil
}

type TileTypeFinder interface {
	TileType(int) TileType // Returns the TileType for the given id
}
