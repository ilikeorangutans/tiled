package tiled

type Point interface {
	X() int
	Y() int
}

func NewPoint(x, y int) Point {
	return &pointImpl{x: x, y: y}
}

type pointImpl struct {
	x, y int
}

func (p *pointImpl) X() int {
	return p.x
}

func (p *pointImpl) Y() int {
	return p.y
}

// Calculates an index within a sequential list of coordinates.
func PointToIndex(p Point, boundaries Rect) int {
	return p.Y()*boundaries.Width() + p.X()
}
