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

func (s *Stack) Len() int {
	return len(s.values)
}

type Interpreter struct {
	Stack Stack
}


