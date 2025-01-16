package queue

type Node struct {
	Value int
	Next *Node
	Prev *Node
}

type Queue struct {
	Head *Node
	Tail *Node
}

func (q *Queue) Enqueue(value int) {
	node := &Node{Value: value}

	if q.Head == nil {
		q.Head = node
	} else {
		q.Tail.Next = node
		node.Prev = q.Tail
	}

	q.Tail = node
}

func (q *Queue) Dequeue() *Node {
	node := q.Head

	if node == nil {
		return nil
	}

	q.Head = node.Next

	if q.Head == nil {
		q.Tail = nil
	} else {
		q.Head.Prev = nil
	}

	return node
}

func (q *Queue) Peek() *Node {
	return q.Head
}

func (q *Queue) Length() int {
	node := q.Head
	length := 0

	for node != nil {
		length++
		node = node.Next
	}

	return length
}