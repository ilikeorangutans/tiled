package tiled

import (
	"fmt"
	"github.com/ilikeorangutans/tiled/tmx"
)

// A single tile on a Map.
type Tile interface {
	Point
	Type() TileType
}

type TileType interface {
	Gid() int
	Tileset() Tileset
}

func TileFromTMX(tmx *tmx.Tile, position Point, tilesets TilesetCatalog) Tile {
	tileType := tilesets.TileType(tmx.Gid)

	return &tmxTile{
		Point:    position,
		tileType: tileType,
	}
}

type tmxTile struct {
	Point
	tileType TileType
}

func (t *tmxTile) Type() TileType {
	return t.tileType
}

type TileTypeFinder interface {
	TileType(int) TileType // Returns the TileType for the given id
}

type tileTypeImpl struct {
	gid     int
	tileset Tileset
}

func (tt *tileTypeImpl) Gid() int {
	return tt.gid
}

func (tt *tileTypeImpl) Tileset() Tileset {
	return tt.tileset
}

func (tt *tileTypeImpl) String() string {
	return fmt.Sprintf("TileType{gid: %d}", tt.gid)
}
