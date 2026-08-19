package interpreter


type Stack struct {
	values []any
}

func (s *Stack) Push(value any) {
	s.values = append(s.values, value)
}

func (s *Stack) Pop() any {
	n := len(s.values)

	value := s.values[n-1]
	s.values = s.values[:n-1]

	return value
}

// Builtin duplicate last item in stack
func (s *Stack) Duplicate() {
	n := len(s.values)
	s.Push(s.values[n-1])
}

// Builtin swap last two items 
func (s *Stack) Swap() {
	n := len(s.values)

	s.values[n-1], s.values[n-2] =
		s.values[n-2], s.values[n-1]
}

// Builtin over method
func (s *Stack) Over() {
	n := len(s.values)

	s.Push(s.values[n-2])
}

// Builtin over stack len 
func (s *Stack) Len() int {
	return len(s.values)
}

