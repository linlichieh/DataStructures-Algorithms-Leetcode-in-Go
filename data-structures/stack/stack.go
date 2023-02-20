package stack

import "log"

type Stack interface {
	Push(string)
	Pop() string
	Top() string
	Size() int
}

func NewStack(typ string) (s Stack) {
	switch typ {
	case "linked_list":
		s = &LinkedListStack{}
	case "slice":
		s = &SliceStack{}
	default:
		log.Fatal(typ, "not supported")
	}
	return
}
