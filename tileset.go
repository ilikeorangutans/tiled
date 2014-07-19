package tiled

import (
	"github.com/ilikeorangutans/tiled/tmx"
)

// A specific tileset
type Tileset interface {
	TileTypeFinder
	Name() string  // Name of the tileset
	FirstGid() int // First GID within this tileset
}

func TilesetFromTmxTileset(t tmx.Tileset) Tileset {

	return &tmxTileset{}
}

type tmxTileset struct {
	tiles    []TileType
	name     string
	firstGid int
}

func (t *tmxTileset) TileType(int) TileType {
	return nil
}

func (t *tmxTileset) FirstGid() int {
	return t.firstGid
}

func (t *tmxTileset) Name() string {
	return t.name
}
