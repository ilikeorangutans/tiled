package tiled

type Rect interface {
	Point
	Width() int
	Height() int
}

func NewRect(x, y, width, height int) Rect {
	return &rectImpl{
		Point:  NewPoint(x, y),
		width:  width,
		height: height,
	}
}

type rectImpl struct {
	Point
	width, height int
}

func (r *rectImpl) Height() int {
	return r.height
}

func (r *rectImpl) Width() int {
	return r.width
}

func Area(r Rect) int {
	return r.Width() * r.Height()
}
