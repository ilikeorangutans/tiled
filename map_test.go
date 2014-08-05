package tiled

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadMap(t *testing.T) {

	m, err := LoadMap("tmx/simple_map_10x10.tmx")
	assert.Nil(t, err)
	assert.NotNil(t, m)

	assert.Equal(t, 10, m.Height())
	assert.Equal(t, 10, m.Width())
	assert.Equal(t, 1, len(m.Layers()))

	layer := m.Layers()[0]
	assert.Equal(t, 0, layer.X())
	assert.Equal(t, 0, layer.Y())

	tile := layer.TileAt(NewPoint(0, 0))
	assert.Equal(t, 0, tile.X())
	assert.Equal(t, 0, tile.Y())

	tile = layer.TileAt(NewPoint(9, 9))
	assert.Equal(t, 9, tile.X())
	assert.Equal(t, 9, tile.Y())
}

func TestSubMap(t *testing.T) {
	m, err := LoadMap("tmx/simple_map_10x10.tmx")
	assert.Nil(t, err)

	sub := m.Sub(NewRect(0, 0, 5, 6))
	assert.NotNil(t, sub)
	assert.Equal(t, 5, sub.Width())
	assert.Equal(t, 6, sub.Height())

	assert.Equal(t, len(m.Layers()), len(sub.Layers()))

	layer := sub.Layers()[0]
	assert.Equal(t, 5, layer.Width())
	assert.Equal(t, 6, layer.Height())

	tile := layer.TileAt(NewPoint(0, 0))
	assert.Equal(t, 0, tile.X())
	assert.Equal(t, 0, tile.Y())
	assert.Equal(t, 1, tile.Type().Gid())

	sub = m.Sub(NewRect(1, 1, 5, 6))
	assert.NotNil(t, sub)
	layer = sub.Layers()[0]
	tile = layer.TileAt(NewPoint(0, 0))
	assert.Equal(t, 1, tile.X())
	assert.Equal(t, 1, tile.Y())
	//assert.Equal(t, 0, tile.Type().Gid())

	sub = m.Sub(NewRect(1, 2, 5, 6))
	assert.NotNil(t, sub)
	assert.Equal(t, 5, sub.Width())
	assert.Equal(t, 6, sub.Height())

	assert.Equal(t, len(m.Layers()), len(sub.Layers()))

	layer = sub.Layers()[0]
	assert.Equal(t, 5, layer.Width())
	assert.Equal(t, 6, layer.Height())

	tile = layer.TileAt(NewPoint(0, 0))
	assert.Equal(t, 1, tile.X())
	assert.Equal(t, 2, tile.Y())
	//assert.Equal(t, 0, tile.Type().Gid())
}

func TestLoadMapTileset(t *testing.T) {
	m, err := LoadMap("tmx/simple_map_10x10.tmx")
	assert.Nil(t, err)

	tilesets := m.Tilesets()
	assert.NotNil(t, tilesets, "Map should have a tileset")
}
