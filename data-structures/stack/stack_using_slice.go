package stack

type SliceStack struct {
	slice  []string
	length int
}

func (s *SliceStack) Push(val string) {
	s.slice = append(s.slice, val)
	s.length++
}

func (s *SliceStack) Pop() string {
	if len(s.slice) == 0 {
		return ""
	}
	popVal := s.slice[len(s.slice)-1]
	s.slice = s.slice[:len(s.slice)-1]
	s.length--
	return popVal
}

func (s *SliceStack) Top() string {
	if len(s.slice) == 0 {
		return ""
	}
	return s.slice[len(s.slice)-1]
}

func (s *SliceStack) Size() int {
	return s.length
}
