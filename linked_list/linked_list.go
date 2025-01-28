package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	FirstNode *Node
}

func NewNode(val int) *Node {
	return &Node{Value: val}
}

func (ll *LinkedList) Add(val int) {
	node := NewNode(val)
	current := ll.FirstNode
	for current.Next != nil {
		current = current.Next
	}

	current.Next = node
}

func (ll *LinkedList) Display() {
	current := ll.FirstNode

	if current != nil {
		for current != nil {
			fmt.Println(current.Value)
			current = current.Next
		}
		//fmt.Println(current.Value)
	} else {
		fmt.Println("Empty Linked List")
	}
}

func (ll *LinkedList) Insert(val, position int) {
	currentPosition := 1
	current := ll.FirstNode

	if current != nil && position > 0 {
		for current.Next != nil && currentPosition != position-1 {
			fmt.Println(current.Value)
			current = current.Next
			currentPosition++
		}

		if currentPosition == position-1 {
			node := NewNode(val)
			current.Next, node.Next = node, current.Next
		} else {
			fmt.Println("Out of range")
		}
	} else {
		fmt.Println("Failed to insert")
	}
}

func (ll *LinkedList) Delete(value int) {
	current := ll.FirstNode
	var prev *Node

	if current != nil {
		for current != nil {
			if current.Value == value {
				if current == ll.FirstNode {
					ll.FirstNode = current.Next
					current.Next = nil
					break
				}
				prev.Next = current.Next
			}
			prev, current = current, current.Next
		}
	}
}

func main() {
	var ll LinkedList

	ll.FirstNode = NewNode(1)

	ll.Add(2)
	ll.Add(3)
	ll.Display()
	fmt.Println("After insertion")
	ll.Insert(4, 2)
	ll.Display()

	fmt.Println("After delete")
	ll.Delete(2)
	ll.Display()

	fmt.Println("Adding more elements")
	ll.Add(5)
	ll.Add(6)
	ll.Add(7)
	ll.Display()

	fmt.Println("Deleting first and last element")
	ll.Delete(1)
	ll.Delete(7)
	ll.Display()

	fmt.Println("Adding more elements")
	ll.Add(8)
	ll.Display()
}
