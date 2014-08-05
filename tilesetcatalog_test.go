package tiled

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindTilesetByTileGid(t *testing.T) {
	tsc := &tilesetCatalog{
		tilesets:  make([]Tileset, 0),
		tileTypes: make(map[int]TileType),
	}

	ts1 := &tmxTileset{
		name:     "Test 1",
		firstGid: 1,
	}
	ts2 := &tmxTileset{
		name:     "Test 2",
		firstGid: 10,
	}
	ts3 := &tmxTileset{
		name:     "Test 3",
		firstGid: 30,
	}

	tsc.Add(ts1)
	tsc.Add(ts2)
	tsc.Add(ts3)

	tileset := tsc.findTilesetByTileGid(1)
	assert.Equal(t, ts1, tileset)

	tileset = tsc.findTilesetByTileGid(9)
	assert.Equal(t, ts1, tileset)

	tileset = tsc.findTilesetByTileGid(10)
	assert.Equal(t, ts2, tileset)

	tileset = tsc.findTilesetByTileGid(31)
	assert.Equal(t, ts3, tileset)
}

func TestTileTypeShouldHaveCorrectTilesetAssociated(t *testing.T) {

	ts1 := &tmxTileset{
		name:     "Test 1",
		firstGid: 1,
	}
	ts2 := &tmxTileset{
		name:     "Test 2",
		firstGid: 10,
	}
	tc := NewTilesetCatalog()

	tc.Add(ts1)
	tc.Add(ts2)

	tileType := tc.TileType(1)
	assert.Equal(t, 1, tileType.Gid())
	assert.Equal(t, ts1, tileType.Tileset())

	tileType = tc.TileType(10)
	assert.Equal(t, 10, tileType.Gid())
	assert.Equal(t, ts2, tileType.Tileset())
}
