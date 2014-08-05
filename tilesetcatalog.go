package tiled

import (
	"sort"
)

// Create a sortable collection of tilesets
type TilesetsByGid []Tileset

func (t TilesetsByGid) Len() int           { return len(t) }
func (t TilesetsByGid) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t TilesetsByGid) Less(i, j int) bool { return t[i].FirstGid() < t[j].FirstGid() }

// Describes a collection of all tilesets
type TilesetCatalog interface {
	Add(Tileset)           // Add the given tileset to the catalog
	TileType(int) TileType // Find the tile type for the given tile gid
}

func NewTilesetCatalog() TilesetCatalog {

	return &tilesetCatalog{
		tilesets:  make([]Tileset, 0),
		tileTypes: make(map[int]TileType),
	}
}

// Holds all available tilesets
type tilesetCatalog struct {
	tilesets  []Tileset
	tileTypes map[int]TileType
}

func (t *tilesetCatalog) TileType(id int) TileType {

	tileType, ok := t.tileTypes[id]
	if !ok {

		tileType = &tileTypeImpl{
			gid:     id,
			tileset: t.findTilesetByTileGid(id),
		}

		t.tileTypes[id] = tileType
	}

	return tileType
}

func (t *tilesetCatalog) Add(ts Tileset) {
	t.tilesets = append(t.tilesets, ts)
	sort.Sort(TilesetsByGid(t.tilesets))
}

// Finds the tileset that would be responsible for the given tile gid.
func (t *tilesetCatalog) findTilesetByTileGid(gid int) Tileset {

	if len(t.tilesets) == 0 {
		return nil
	}

	lastTileset := t.tilesets[0]

	for key := range t.tilesets {

		cur := t.tilesets[key]

		if gid >= lastTileset.FirstGid() && gid < cur.FirstGid() {
			return lastTileset
		}

		lastTileset = cur
	}

	return lastTileset
}
