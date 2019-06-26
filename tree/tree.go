package main

import (
	"fmt"
)

type Stack struct {
	elements []*TreeNode
}

type TreeNode struct {
	key       int
	leftNode  *TreeNode
	rightNode *TreeNode
}

type BinarySearchTree struct {
	rootNode *TreeNode
}

func (s *Stack)Push(item *TreeNode) {
	s.elements = append(s.elements, item)
}

func (s *Stack)Pop() *TreeNode {

	if s.IsEmpty() {
		return nil
	}

	len := len(s.elements)
	ele := s.elements[len-1]
	s.elements = s.elements[:len-1]

	return ele
}

func (s *Stack)IsEmpty() bool {
	return len(s.elements) == 0 
}

func insertNode(rootNode *TreeNode, newNode *TreeNode) {
	if newNode.key < rootNode.key {
		if rootNode.leftNode == nil {
			rootNode.leftNode = newNode
		} else {
			insertNode(rootNode.leftNode, newNode)
		}
	} else {
		if rootNode.rightNode == nil {
			rootNode.rightNode = newNode
		} else {
			insertNode(rootNode.rightNode, newNode)
		}
	}
}

func (bst *BinarySearchTree)InsertElement(key int) {
	newNode := &TreeNode{key, nil, nil}
	if bst.rootNode == nil {
		bst.rootNode = newNode
		return
	} else {
		insertNode(bst.rootNode, newNode)
	}
}

func (bst *BinarySearchTree)PreOrderTraversal() {
	s := &Stack{}

	root := bst.rootNode

	for true {
			for root != nil {
				fmt.Println(root.key)
				s.Push(root)
				root = root.leftNode
			}

			if s.IsEmpty() {
				break
			}
			root = s.Pop()
			root = root.rightNode
			
	}
}
func main() {
	bst := &BinarySearchTree{}
	bst.InsertElement(10)
	bst.InsertElement(9)
	bst.InsertElement(11)
	bst.InsertElement(8)
	bst.InsertElement(15)
	bst.InsertElement(12)
	bst.InsertElement(5)
	bst.InsertElement(7)
	bst.InsertElement(6)
	bst.PreOrderTraversal()
	
}