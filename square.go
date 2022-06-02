package square

type Point struct {
	x, y int
}
type Square struct {
	P Point
	a uint
}

func (s *Square) endPoint() Point {
	return Point{
		x: s.P.x + int(s.a),
		y: s.P.y + int(s.a),
	}
}
func (s *Square) Area() uint {
	return s.a * s.a
}

func (s *Square) Perimeter() uint {
	return 4 * s.a
}
