package main

import (
	"fmt"
)

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
	if parent.Left == nil {
		parent.Left = NewNode(value)
	} else {
		newNode := NewNode(value)
		parent.Left, newNode.Left = newNode, parent.Left
	}
}

func (bt *BinaryTree) InsertRight(parent *Node, value int) {
	if parent.Right == nil {
		parent.Right = NewNode(value)
	} else {
		newNode := NewNode(value)
		parent.Right, newNode.Left = newNode, parent.Right
	}
}

func (bt *BinaryTree) InOrder(parent *Node) {
	stack := []*Node{}
	current := parent

	finalOp := []int{}

	for current != nil || len(stack) > 0 {
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		finalOp = append(finalOp, current.Value)

		current = current.Right
	}

	fmt.Println("InOrder: ", finalOp)
}

func (bt *BinaryTree) PreOrder(parent *Node) {
	stack := []*Node{}
	current := parent
	finalOp := []int{}

	stack = append(stack, current)
	for len(stack) > 0 {
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		finalOp = append(finalOp, current.Value)

		if current.Right != nil {
			stack = append(stack, current.Right)
		}

		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}

	fmt.Println("PreOrder: ", finalOp)
}

func (bt *BinaryTree) PostOrder(root *Node) {
	var stack1 []*Node
	var stack2 []*Node
	finalOp := []int{}

	if root != nil {
		stack1 = append(stack1, root)

		for len(stack1) > 0 {
			node := stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]

			stack2 = append(stack2, node)

			if node.Left != nil {
				stack1 = append(stack1, node.Left)
			}

			if node.Right != nil {
				stack1 = append(stack1, node.Right)
			}
		}

		for len(stack2) > 0 {
			node := stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
			finalOp = append(finalOp, node.Value)
		}
	}

	fmt.Println("PostOrder: ", finalOp)
}

func main() {

	bt := BinaryTree{
		Root: NewNode(1),
	}

	bt.InsertLeft(bt.Root, 2)
	bt.InsertRight(bt.Root, 3)

	bt.InsertLeft(bt.Root.Left, 4)
	bt.InsertRight(bt.Root.Left, 5)

	bt.InsertLeft(bt.Root.Right, 6)
	bt.InsertRight(bt.Root.Right, 7)

	fmt.Println("====PreOrder=====")
	bt.PreOrder(bt.Root)
	fmt.Println("====InOrder=====")
	bt.InOrder(bt.Root)
	fmt.Println("====PostOrder=====")
	bt.PostOrder(bt.Root)
}
