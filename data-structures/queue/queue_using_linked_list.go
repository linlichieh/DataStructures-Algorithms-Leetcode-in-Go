package queue

type node struct {
	val  string
	next *node
}

type LinkedListQueue struct {
	first  *node
	last   *node
	length int
}

func (q *LinkedListQueue) Enqueue(val string) {
	newNode := &node{val: val}
	if q.length == 0 {
		q.first = newNode
		q.last = newNode
	} else {
		q.last.next = newNode
		q.last = newNode
	}
	q.length++
}

func (q *LinkedListQueue) Dequeue() string {
	if q.length == 0 {
		return ""
	}
	if q.first == q.last {
		q.last = nil
	}
	holdingNode := q.first
	q.first = q.first.next
	q.length--
	return holdingNode.val
}

func (q *LinkedListQueue) Peek() string {
	if q.first != nil {
		return q.first.val
	}
	return ""
}

func (q *LinkedListQueue) Size() int {
	return q.length
}
