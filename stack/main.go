package stack

type Node struct {
	Value int
	Next *Node
	Prev *Node
}

type Stack struct {
	Head *Node
	Tail *Node
}

func (s *Stack) Push(value int) {
	node := &Node{Value: value}

	if s.Head == nil {
		s.Head = node
	} else {
		s.Tail.Next = node
		node.Prev = s.Tail
	}

	s.Tail = node
}

func (s *Stack) Pop() *Node {
	node := s.Tail

	if node == nil {
		return nil
	}

	s.Tail = node.Prev

	if s.Tail == nil {
		s.Head = nil
	} else {
		s.Tail.Next = nil
	}

	return node
}

func (s *Stack) Peek() *Node {
	return s.Tail
}