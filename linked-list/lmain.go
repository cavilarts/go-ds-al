package linkedlist

type Node struct {
	Value int
	Next *Node
	Prev *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (l *LinkedList) Add(value int) {
	node := &Node{Value: value}

	if l.Head == nil {
		l.Head = node
	} else {
		l.Tail.Next = node
		node.Prev = l.Tail
	}

	l.Tail = node
}

func (l *LinkedList) Get(index int) *Node {
	node := l.Head

	for i := 0; i < index; i++ {
		node = node.Next
	}

	return node
}

func (l *LinkedList) Remove(index int) {
	node := l.Get(index)

	if node.Prev == nil {
		l.Head = node.Next
	} else {
		node.Prev.Next = node.Next
	}

	if node.Next == nil {
		l.Tail = node.Prev
	} else {
		node.Next.Prev = node.Prev
	}
}

func (l *LinkedList) Length() int {
	node := l.Head
	length := 0

	for node != nil {
		length++
		node = node.Next
	}

	return length
}

func (l *LinkedList) Traverse() []int {
	node := l.Head
	values := []int{}

	for node != nil {
		values = append(values, node.Value)
		node = node.Next
	}

	return values
}