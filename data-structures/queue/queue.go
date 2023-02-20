package queue

import "log"

type Queue interface {
	Enqueue(string)
	Dequeue() string
	Peek() string
	Size() int
}

func NewQueue(typ string) (q Queue) {
	switch typ {
	case "linked_list":
		q = &LinkedListQueue{}
	default:
		log.Fatal(typ, "not supported")
	}
	return
}
