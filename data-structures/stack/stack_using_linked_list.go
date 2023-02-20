package stack

type node struct {
	val  string
	next *node
}

type LinkedListStack struct {
	top    *node
	bottom *node
	length int
}

func (s *LinkedListStack) Push(val string) {
	newNode := &node{val: val}
	if s.length == 0 {
		s.top = newNode
		s.bottom = newNode
	} else {
		newNode.next = s.top
		s.top = newNode
	}
	s.length++
}

func (s *LinkedListStack) Pop() string {
	if s.top == nil {
		return ""
	}
	if s.top == s.bottom {
		s.bottom = nil
	}
	holdingNode := s.top
	s.top = holdingNode.next
	s.length--
	return holdingNode.val
}

func (s *LinkedListStack) Top() string {
	if s.top != nil {
		return s.top.val
	}
	return ""
}

func (s *LinkedListStack) Size() int {
	return s.length
}
