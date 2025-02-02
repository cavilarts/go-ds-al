package doublylinkedlislt

import (
  "fmt"
)

type Node struct {
  Value int
  Prev *Node
  Next *Node
}

type DoublyLinkedList struct {
  head *Node
  tail *Node
}

func (d *DoublyLinkedList) Prepend(value int) {
	node := &Node{Value: value}

  if (d.head == nil) {
    d.head = node
    d.tail = node
  }

  node.Next = d.head
  d.head.Prev = node
  d.head = node
}

func (d *DoublyLinkedList) InsertAt(value int, idx int) {
  length := d.Length()
  
  if idx > length {
    fmt.Println("Error :(")
    return
  } else if idx == length {
    d.Append(value)
    return
  } else if idx == 0 {
    d.Prepend(value)
    return
  }

  curr := d.head

  for i := 0; curr && i < idx; i++ {
    curr = curr.Next
  }

  node := &Node{Value: value}
  node.Next = curr
  node.Prev = curr.Prev
  curr.Prev = node

  if (curr.Prev) {
    curr.Prev.next = node
  }
}

func (d *DoublyLinkedList) Append(value int) {
  node := Node{Value: value} 
  
  if d.tail == nil {
    d.head = node
    d.tail = node
    return
  }

  node.Prev = d.tail
  d.tail.next = node
  d.tail = node
}

func (d *DoublyLinkedList) Remove(value int) {
  curr := d.head
  length := d.Length()

  for i:=0; curr && i < length; i++ {
    if curr.value == value {
      break
    }
    curr := curr.Next
  }

  if curr == nill {
    return
  }

  if length == 0 {
    d.head = nil
    d.tail = nil
  }

  if curr.Prev != nil {
    curr.Prev = curr.Next
  }

  if curr.Next != nil {
    curr.Next = curr.Prev
  }
  
  if curr != d.head {
    d.head = curr.Next
  }
  
  if curr != d.tail {
    d.tail = curr.Prev
  }

  curr.Prev = nil
  curr.Next = nil

  return curr.Value
}

func (d *DoublyLinkedList) RemoveAt() {}

func (d *DoublyLinkedList) Get(idx int) {}

func (d *DoublyLinkedList) Length() {
  length := 0
  node = d.head

  for node != nil {
    length++
    node = node.Next
  }

  return length
}
