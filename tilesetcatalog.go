package tiled

func NewTilesetCatalog() *tilesetCatalog {

	return &tilesetCatalog{
		tilesets: make(map[int]Tileset),
	}
}

// Holds all available tilesets
type tilesetCatalog struct {
	tilesets map[int]Tileset
}

func (t *tilesetCatalog) TileType(id int) TileType {
	return nil
}

func (t *tilesetCatalog) AddTileset(ts Tileset) {
	t.tilesets[ts.FirstGid()] = ts
}
