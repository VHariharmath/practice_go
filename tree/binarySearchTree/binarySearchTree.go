package main

import "fmt"

type treeNode struct {
	key       int
	leftNode  *treeNode
	rightNode *treeNode
}

type binarySearchTree struct {
	rootNode *treeNode
}

type Stack struct {
	elements []*treeNode
}

func (s *Stack) Push(item *treeNode) {
	s.elements = append(s.elements, item)
}

func (s *Stack) Pop() *treeNode {

	if s.IsEmptyStack() {
		return nil
	}

	len := len(s.elements)
	ele := s.elements[len-1]
	s.elements = s.elements[:len-1]

	return ele
}

func (s *Stack) IsEmptyStack() bool {
	return len(s.elements) == 0
}

func CreateStack() *Stack {
	s := &Stack{}
	s.elements = make([]*treeNode, 0)
	return s
}

func (bst *binarySearchTree) insertElement(item int) {
	newNode := &treeNode{item, nil, nil}
	root := bst.rootNode
	if root == nil {
		bst.rootNode = newNode
		return
	} else {
		insertNode(root, newNode)
	}

}

func insertNode(root *treeNode, newNode *treeNode) {

	if newNode.key < root.key {
		if root.leftNode != nil {
			insertNode(root.leftNode, newNode)
		} else {
			root.leftNode = newNode
		}
	} else {
		if root.rightNode != nil {
			insertNode(root.rightNode, newNode)
		} else {
			root.rightNode = newNode
		}
	}
}

func (bst *binarySearchTree) InOrderTraversal() {
	root := bst.rootNode
	if root == nil {
		return
	}

	s := CreateStack()

	for true {
		for root != nil {
			s.Push(root)
			root = root.leftNode
		}

		if s.IsEmptyStack() {
			break
		}

		root = s.Pop()
		fmt.Println(root.key)
		root = root.rightNode
	}

}

func (bst *binarySearchTree) findAnElement(key int) bool {
	root := bst.rootNode
	if root == nil {
		return false
	}

	for root != nil {
		if root.key == key {
			return true
		} else if key < root.key {
			root = root.leftNode
		} else {
			root = root.rightNode
		}
	}

	return false
}

func (bst *binarySearchTree) findMinimum() int {
	root := bst.rootNode
	if root == nil {
		return 0
	}

	for root.leftNode != nil {
		root = root.leftNode
	}

	return root.key
}

func (bst *binarySearchTree) findMax() int {
	root := bst.rootNode
	if root == nil {
		return 0
	}

	for root.rightNode != nil {
		root = root.rightNode
	}
	return root.key
}

func (bst *binarySearchTree) predecessor() int {
	root := bst.rootNode
	if root == nil {
		return 0
	}

	var prev *treeNode
	for root.leftNode != nil {
		prev = root
		root = root.leftNode
	}

	
}

func main() {
	bst := &binarySearchTree{}
	bst.insertElement(10)
	bst.insertElement(9)
	bst.insertElement(11)
	bst.insertElement(8)
	bst.insertElement(15)
	bst.insertElement(12)
	bst.insertElement(5)
	bst.insertElement(7)
	bst.insertElement(6)
	bst.insertElement(1)
	//bst.InOrderTraversal()

	//fmt.Println(bst.findAnElement(12))
	//fmt.Println(bst.findAnElement(123))

	fmt.Println(bst.findMinimum())
	fmt.Println(bst.findMax())
}
