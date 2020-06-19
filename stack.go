package main

// Stack ...Only supports interface type. Primitives not supported.
type Stack struct {
	elements []interface{}
}

// NewStack ...
func NewStack() *Stack {
	return &Stack{}
}

// remove the first element
func (s *Stack) pop() (interface{}, bool) {
	if len(s.elements) != 0 {
		temp := s.elements[0]
		copy(s.elements[0:], s.elements[1:])
		s.elements[len(s.elements)-1] = nil
		s.elements = s.elements[:len(s.elements)-1]
		return temp, true
	}

	return nil, false
}

// Add to the top
func (s *Stack) push(elem interface{}) {
	s.elements = append(s.elements, elem)
	copy(s.elements[1:], s.elements)
	s.elements[0] = elem
}

func (s *Stack) len() int {
	return len(s.elements)
}
