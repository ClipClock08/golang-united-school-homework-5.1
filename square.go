package golang_united_school_homework_5_1

type Square struct {
	A int
	B int
}

func (s *Square) Area() int {
	return s.A * s.B
}

func (s *Square) Perimetr() int {
	return 2 * (s.A + s.B)
}
