package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func NewNode(value int) *Node {
	return &Node{Value: value}
}

type BinaryTree struct {
	Root *Node
}

func (bt *BinaryTree) InsertLeft(parent *Node, value int) {
	if parent == nil {
		parent.Left = NewNode(value)
	} else {
		newNode := NewNode(value)
		newNode.Left, parent.Left = parent.Left, newNode
	}
}

func (bt *BinaryTree) InsertRight(parent *Node, value int) {
	if parent == nil {
		parent.Right = NewNode(value)
	} else {
		newNode := NewNode(value)
		newNode.Right, parent.Right = parent.Right, newNode
	}
}

func (bt *BinaryTree) InOrder(parent *Node) {
	if parent != nil {
		bt.InOrder(parent.Left)
		fmt.Println("Value: ", parent.Value)
		bt.InOrder(parent.Right)
	}
}

func (bt *BinaryTree) PreOrder(parent *Node) {
	if parent != nil {
		fmt.Println("Value: ", parent.Value)
		bt.PreOrder(parent.Left)
		bt.PreOrder(parent.Right)
	}
}

func (bt *BinaryTree) PostOrder(parent *Node) {
	if parent != nil {
		bt.PostOrder(parent.Left)
		bt.PostOrder(parent.Right)
		fmt.Println("Value: ", parent.Value)
	}
}

func main() {
	bt := &BinaryTree{
		Root: NewNode(1),
	}

	bt.InsertLeft(bt.Root, 2)
	bt.InsertRight(bt.Root, 3)

	bt.InsertLeft(bt.Root.Left, 4)
	bt.InsertRight(bt.Root.Left, 5)

	bt.InsertLeft(bt.Root.Right, 6)
	bt.InsertRight(bt.Root.Right, 7)

	fmt.Println("------InOrder----------")
	bt.InOrder(bt.Root)

	fmt.Println("------PreOrder----------")
	bt.PreOrder(bt.Root)

	fmt.Println("------PostOrder----------")
	bt.PostOrder(bt.Root)
}
