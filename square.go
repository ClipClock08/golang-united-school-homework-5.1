package square

type Square struct {
	A uint
}

func (s *Square) Area() uint {
	return s.A * s.A
}

func (s *Square) Perimeter() uint {
	return 4 * s.A
}
