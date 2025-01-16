package main

import (
	linkedlist "carlosavila.com/go/linked-list"
)

func main() {
	linked := linkedlist.LinkedList{}
	linked.Add(1)
	linked.Add(2)
	linked.Add(3)

	linked.Remove(1)

	node := linked.Get(1)
	println(node.Value)

	println(linked.Length())

	linked.Add(4)

	println(linked.Length())

	linked.Remove(0)

	println(linked.Length())

	linked.Traverse()
}