package viewport

// Describes a viewport on an area. The viewport is usually smaller than the area
// but that's not a necessity. The viewport has a position relative to the origin
// (top left corner with coordinates 0/0) of the area and a width and height.
type Viewport interface {
	// Query methods:

	X() int          // Offset in x-axis on the area
	Y() int          // Offset in y-axis on the area
	Width() int      // width of the viewport in pixels
	Height() int     // height of the viewport in pixels
	AreaWidth() int  // width of the overall area this viewport is showing in pixels
	AreaHeight() int // height of the area in pixels

	// Movement methods

	MoveBy(deltaX, deltaY int) // Move the viewport relative by the given deltas
	MoveTo(x, y int)           // Move the viewport to this absolute position

	// Viewport calculations:

	VisibleTiles() (int, int, int, int) // Returns the top left and bottom right coordinates of the rectangle of visible tiles
}

func NewViewport(width, height, x, y, widthInTiles, heightInTiles int) Viewport {
	tileWidth := 32
	tileHeight := 32

	areaWidth := widthInTiles * tileWidth
	areaHeight := heightInTiles * tileHeight
	return &viewportImpl{
		x:             x,
		y:             y,
		width:         width,
		height:        height,
		areaWidth:     areaWidth,
		areaHeight:    areaHeight,
		maxX:          (areaWidth - width) - tileWidth,
		maxY:          (areaHeight - height) - tileHeight,
		tileWidth:     tileWidth,
		tileHeight:    tileHeight,
		widthInTiles:  widthInTiles,  // width of the area in tiles
		heightInTiles: heightInTiles, // height of the are in tiles
		visibleTilesX: width>>5 + 2,  // number of tiles visible on the x-axis
		visibleTilesY: height>>5 + 2, // number of tiles visible on the y-axis
	}
}

type viewportImpl struct {
	x, y                         int // x and y coordinate
	maxX, maxY                   int // maximum x and y
	areaWidth, areaHeight        int
	width, height                int
	tileWidth, tileHeight        int
	widthInTiles, heightInTiles  int
	visibleTilesX, visibleTilesY int // Number of visible tiles in x and y direction
}

func (v *viewportImpl) X() int {
	return v.x
}

func (v *viewportImpl) Y() int {
	return v.y
}

func (v *viewportImpl) AreaWidth() int {
	return v.areaWidth
}

func (v *viewportImpl) AreaHeight() int {
	return v.areaHeight
}

func (v *viewportImpl) Width() int {
	return v.width
}

func (v *viewportImpl) Height() int {
	return v.height
}

func (v *viewportImpl) MoveBy(deltaX, deltaY int) {
	v.MoveTo(v.x+deltaX, v.y+deltaY)
}

func (v *viewportImpl) MoveTo(x, y int) {
	v.x = bounded(0, v.maxX, x)
	v.y = bounded(0, v.maxY, y)
}

// Calculates the rectangle of the currently visible tiles
func (v *viewportImpl) VisibleTiles() (minX, minY, maxX, maxY int) {
	minX = v.X() >> 5
	minY = v.Y() >> 5
	maxX = minInt(minX+v.visibleTilesX, v.widthInTiles-1)
	maxY = minInt(minY+v.visibleTilesY, v.heightInTiles-1)
	return minX, minY, maxX, maxY
}

// Returns a the value if it is within the given lower or upper boundary. If the
// value is smaller or larger than the lower or upper value, the respective
// boundary value will be returned.
func bounded(lower, upper, value int) int {
	if value < lower {
		return lower
	} else if value > upper {
		return upper
	} else {
		return value
	}
}

// Returns the smaller of the two ints
func minInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}

}
