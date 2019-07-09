package main

import "fmt"

type avlTreeNode struct {
	key    int
	height int
	left   *avlTreeNode
	right  *avlTreeNode
}

type avlTreeHead struct {
	root *avlTreeNode
}

func height(root *avlTreeNode) int {
	if root == nil {
		return -1
	} else {
		return root.height
	}
}

func max(h1 int, h2 int) int {
	if h1 > h2 {
		return h1
	} else {
		return h2
	}
}

func singleRotateRight(x *avlTreeNode) *avlTreeNode {
	w := x.left
	x.left = w.right
	w.right = x

	x.height = max(height(x.left), height(x.right)) + 1
	w.height = max(height(w.left), height(w.right)) + 1
	return w
}

func singleRotateLeft(x *avlTreeNode) *avlTreeNode {
	w := x.right
	x.right = w.left
	w.left = x

	x.height = max(height(x.left), height(x.right)) + 1
	w.height = max(height(w.left), height(w.right)) + 1

	return w
}

func getBalanceFactor(x *avlTreeNode) int {
	if x == nil {
		return 0
	}

	return height(x.left) - height(x.right)
}

func doubleRotateLeftRight(x *avlTreeNode) *avlTreeNode {

	x.left = singleRotateLeft(x.left)
	return singleRotateRight(x)
}

func doubleRotateRightLeft(x *avlTreeNode) *avlTreeNode {
	x.right = singleRotateRight(x.right)

	return singleRotateLeft(x)
}

func (rootHead *avlTreeHead) insert(key int) {

	newNode := &avlTreeNode{key, 0, nil, nil}
	root := rootHead.root
	if root == nil {
		rootHead.root = newNode
		return
	} else {
		insertNode(root, newNode)
	}
}

func insertNode(root *avlTreeNode, newNode *avlTreeNode) {

	if newNode.key < root.key {
		if root.left == nil {
			root.left = newNode
		} else {
			insertNode(root.left, newNode)
		}
	}

	if newNode.key > root.key {
		if root.right == nil {
			root.right = newNode
		} else {
			insertNode(root.right, newNode)
		}
	}

	root.height = max(height(root.left), height(root.right)) + 1

	balance := getBalanceFactor(root)

	if balance > 1 && newNode.key < root.left.key {
		root = singleRotateRight(root)
	}

	if balance < -1 && newNode.key > root.right.key {
		root = singleRotateLeft(root)
	}

	if balance > 1 && newNode.key > root.left.key {
		root = doubleRotateLeftRight(root)
	}

	if balance < -1 && newNode.key < root.right.key {
		root = doubleRotateRightLeft(root)
	}
}

type queue struct {
	elements []*avlTreeNode
}

func createQueue() queue {
	var q queue
	q.elements = make([]*avlTreeNode, 0)
	return q
}

func (q *queue) enqueue(node *avlTreeNode) {
	q.elements = append(q.elements, node)
}

func (q *queue) dequeue() *avlTreeNode {

	if q.isEmptyQueue() {
		return nil
	}
	temp := q.elements[0]
	q.elements = q.elements[1:len(q.elements)]
	return temp
}

func (q *queue) isEmptyQueue() bool {

	return len(q.elements) == 0
}

func (rootHead *avlTreeHead) levelOrderTraversal() {
	root := rootHead.root
	if root == nil {
		return
	}

	q := createQueue()
	q.enqueue(root)

	for !q.isEmptyQueue() {
		temp := q.dequeue()

		fmt.Println(temp.key)
		if temp.left != nil {
			q.enqueue(temp.left)
		}

		if temp.right != nil {
			q.enqueue(temp.right)
		}

	}
}

func main() {
	avl := &avlTreeHead{}
	avl.insert(10)
	avl.insert(9)
	avl.insert(11)
	avl.insert(8)
	avl.insert(15)
	avl.insert(12)
	avl.insert(5)
	avl.insert(7)
	avl.insert(6)
	avl.insert(1)
	avl.levelOrderTraversal()
}
