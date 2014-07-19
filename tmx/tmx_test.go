package tmx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadFilenameSimpleMap(t *testing.T) {
	m, err := LoadTmxWithFilename("simple_map_10x10.tmx")

	assert.Nil(t, err)
	assert.NotNil(t, m)

	assert.Equal(t, 10, m.Height)
	assert.Equal(t, 10, m.Width)
	assert.Equal(t, 32, m.Tilewidth)
	assert.Equal(t, 32, m.Tileheight)
	assert.Equal(t, "1.0", m.Version)
	assert.Equal(t, "orthogonal", m.Orientation)

	assert.Equal(t, 1, len(m.Tileset))

	tileset := m.Tileset[0]
	assert.Equal(t, 1, tileset.Firstgid)
	assert.Equal(t, 0, tileset.Margin)
	assert.Equal(t, "Simple Tileset", tileset.Name)
	assert.Equal(t, "", tileset.Source)
	assert.Equal(t, 0, tileset.Spacing)
	assert.Equal(t, 32, tileset.Tileheight)
	assert.Equal(t, 32, tileset.Tilewidth)

	assert.Equal(t, "simple_tileset.png", tileset.Image.Source)
	assert.Equal(t, 128, tileset.Image.Width)
	assert.Equal(t, 32, tileset.Image.Height)

	assert.Equal(t, 1, len(m.Layer))
	layer := m.Layer[0]
	assert.Equal(t, "Tile Layer 1", layer.Name)
	assert.Equal(t, 100, len(layer.Data.Tiles))
}

func TestLoadFilenameMultiTilesetMap(t *testing.T) {
	m, err := LoadTmxWithFilename("simple_map_10x10_multitileset.tmx")

	assert.Nil(t, err)

	assert.Equal(t, 2, len(m.Tileset))
}
