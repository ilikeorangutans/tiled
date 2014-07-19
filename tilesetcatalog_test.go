package tiled

import (
	"testing"
)

func TestFoo(t *testing.T) {

	ts := &tmxTileset{
		name:     "Test",
		firstGid: 1,
	}
	tc := NewTilesetCatalog()

	tc.AddTileset(ts)

}
