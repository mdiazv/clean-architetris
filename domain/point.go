package domain

func (p Point) plus(o Point) Point {
	return Point{p.X + o.X, p.Y + o.Y}
}
