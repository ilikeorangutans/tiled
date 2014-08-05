package tiled

import (
	"github.com/ilikeorangutans/tiled/tmx"
)

func LoadMap(filename string) (Map, error) {
	tmxMap, err := tmx.LoadTmxWithFilename(filename)
	if err != nil {
		return nil, err
	}

	return MapFromTMX(tmxMap), nil
}

// Creates a new Map from the given TMX data
func MapFromTMX(m *tmx.Map) Map {

	tilesets := NewTilesetCatalog()

	boundaries := NewRect(0, 0, m.Width, m.Height)

	layers := make([]Layer, len(m.Layer))

	for i := range m.Layer {
		tmxLayer := m.Layer[i]
		layers[i] = LayerFromTMX(boundaries, tilesets, tmxLayer)
	}

	return &tmxMap{
		Rect:     boundaries,
		layers:   layers,
		tilesets: tilesets,
	}
}
