package main

import (
	"fmt"
)

type Stack struct {
	elements []*TreeNode
}

type Queue struct {
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

func (s *Stack) Push(item *TreeNode) {
	s.elements = append(s.elements, item)
}

func (s *Stack) Pop() *TreeNode {

	if s.IsEmpty() {
		return nil
	}

	len := len(s.elements)
	ele := s.elements[len-1]
	s.elements = s.elements[:len-1]

	return ele
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (q *Queue) Enqueue(item *TreeNode) {
	q.elements = append(q.elements, item)

}

func (q *Queue) Dequeue() *TreeNode {
	if q.IsEmptyQueue() {
		return nil
	}

	ele := q.elements[0]
	q.elements = q.elements[1:len(q.elements)]

	return ele
}

func (q *Queue) IsEmptyQueue() bool {
	return len(q.elements) == 0
}

func CreateQueue() *Queue {
	q := &Queue{}
	q.elements = make([]*TreeNode, 0)
	return q
}

func CreateStack() *Stack {
	s := &Stack{}
	s.elements = make([]*TreeNode, 0)
	return s
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

func (bst *BinarySearchTree) InsertElement(key int) {
	newNode := &TreeNode{key, nil, nil}
	if bst.rootNode == nil {
		bst.rootNode = newNode
		return
	} else {
		insertNode(bst.rootNode, newNode)
	}
}

func (bst *BinarySearchTree) PreOrderTraversal() {
	s := CreateStack()

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

func (bst *BinarySearchTree) InOrderTraversal() {
	s := CreateStack()
	root := bst.rootNode

	for true {
		for root != nil {
			s.Push(root)
			root = root.leftNode
		}

		if s.IsEmpty() {
			break
		}

		root = s.Pop()
		fmt.Println(root.key)
		root = root.rightNode
	}
}

func (bst *BinarySearchTree) PostOrderTraversal() {
	s := CreateStack()
	root := bst.rootNode

	s.Push(root)
	var prev *TreeNode

	for !s.IsEmpty() {
		cur := s.Pop()

		if prev == nil || prev.leftNode == cur || prev.rightNode == cur {
			if cur.leftNode != nil {
				s.Push(cur.leftNode)
			} else if cur.rightNode != nil {
				s.Push(cur.rightNode)
			}
		} else if cur.leftNode == prev {
			if cur.rightNode != nil {
				s.Push(cur.rightNode)
			}
		} else {
			fmt.Println(cur.key)
			s.Pop()
		}
		prev = cur
	}
}

func (bst *BinarySearchTree) LevelOrderTraversal() {
	q := CreateQueue()

	root := bst.rootNode
	if root == nil {
		return
	}
	q.Enqueue(root)

	for !q.IsEmptyQueue() {
		temp := q.Dequeue()
		fmt.Println(temp.key)
		if temp.leftNode != nil {
			q.Enqueue(temp.leftNode)
		}

		if temp.rightNode != nil {
			q.Enqueue(temp.rightNode)
		}
	}
}

func (bst *BinarySearchTree) FindMax() int {
	q := CreateQueue()

	root := bst.rootNode
	if root == nil {
		return 0
	}

	max := root.key
	q.Enqueue(root)

	for !q.IsEmptyQueue() {
		temp := q.Dequeue()

		if max < temp.key {
			max = temp.key
		}

		if temp.leftNode != nil {
			q.Enqueue(temp.leftNode)
		}

		if temp.rightNode != nil {
			q.Enqueue(temp.rightNode)
		}
	}

	return max
}

func (bst *BinarySearchTree) SearchAnElement(key int) bool {
	q := CreateQueue()

	root := bst.rootNode
	if root == nil {
		return false
	}
	q.Enqueue(root)

	for !q.IsEmptyQueue() {
		temp := q.Dequeue()
		if temp.key == key {
			return true
		}
		if temp.leftNode != nil {
			q.Enqueue(temp.leftNode)
		}

		if temp.rightNode != nil {
			q.Enqueue(temp.rightNode)
		}
	}

	return false
}

/*
func (bst *BinarySearchTree) InsertElement(key int) {
	newNode := &TreeNode{}
	newNode.key = key

	root := bst.rootNode
	if root == nil {
		bst.rootNode = newNode
		return
	}

	q := CreateQueue()
	q.Enqueue(root)

	for !q.IsEmptyQueue() {
		temp := q.Dequeue()
		if temp.leftNode != nil {
			q.Enqueue(temp.leftNode)
		} else {
			temp.leftNode = newNode
			break
		}

		if temp.rightNode != nil {
			q.Enqueue(temp.rightNode)
		} else {
			temp.rightNode = newNode
			break
		}
	}
}
*/

func (bst *BinarySearchTree) SizeofBinaryTree() int {

	root := bst.rootNode
	if root == nil {
		return 0
	}

	q := CreateQueue()
	count := 0

	q.Enqueue(root)

	for !q.IsEmptyQueue() {
		temp := q.Dequeue()
		count++
		if temp.leftNode != nil {
			q.Enqueue(temp.leftNode)
		}

		if temp.rightNode != nil {
			q.Enqueue(temp.rightNode)
		}
	}

	return count
}

func (bst *BinarySearchTree) PrintLevelOrderinReverse() {

	root := bst.rootNode
	if root == nil {
		return
	}

	q := CreateQueue()
	s := CreateStack()

	q.Enqueue(root)

	for !q.IsEmptyQueue() {
		temp := q.Dequeue()

		s.Push(temp)

		if temp.leftNode != nil {
			q.Enqueue(temp.leftNode)
		}

		if temp.rightNode != nil {
			q.Enqueue(temp.rightNode)
		}
	}

	for !s.IsEmpty() {
		fmt.Println(s.Pop().key)
	}

}

func (bst *BinarySearchTree) LevelsofBinaryTree() int {

	root := bst.rootNode
	if root == nil {
		return 0
	}

	q := CreateQueue()
	level := 0

	q.Enqueue(root)
	q.Enqueue(nil)

	for !q.IsEmptyQueue() {
		temp := q.Dequeue()

		if temp == nil {
			if !q.IsEmptyQueue() {
				q.Enqueue(nil)
			}
			level++
		} else {
			if temp.leftNode != nil {
				q.Enqueue(temp.leftNode)
			}

			if temp.rightNode != nil {
				q.Enqueue(temp.rightNode)
			}
		}
	}

	return level
}

func (bst *BinarySearchTree) lastNodeInBT() *TreeNode {
	q := CreateQueue()
	temp := &TreeNode{}
	root := bst.rootNode

	if root == nil {
		return nil
	}
	q.Enqueue(root)

	for !q.IsEmptyQueue() {
		temp = q.Dequeue()
		if temp.leftNode != nil {
			q.Enqueue(temp.leftNode)
		}

		if temp.rightNode != nil {
			q.Enqueue(temp.rightNode)
		}
	}

	return temp
}

func (bst *BinarySearchTree) NumberOfFullNodes() int {
	root := bst.rootNode
	if root == nil {
		return 0
	}

	q := CreateQueue()
	q.Enqueue(root)
	count := 0

	for !q.IsEmptyQueue() {
		temp := q.Dequeue()

		if temp.leftNode != nil && temp.rightNode != nil {
			count++
		}

		if temp.leftNode != nil {
			q.Enqueue(temp.leftNode)
		}

		if temp.rightNode != nil {
			q.Enqueue(temp.rightNode)
		}
	}

	return count
}

func (bst *BinarySearchTree) NumberOfHalfNodes() int {
	root := bst.rootNode
	if root == nil {
		return 0
	}

	q := CreateQueue()
	q.Enqueue(root)
	count := 0

	for !q.IsEmptyQueue() {
		temp := q.Dequeue()

		if (temp.rightNode == nil && temp.leftNode != nil) || (temp.leftNode == nil && temp.rightNode != nil) {
			count++
		}

		if temp.leftNode != nil {
			q.Enqueue(temp.leftNode)
		}

		if temp.rightNode != nil {
			q.Enqueue(temp.rightNode)
		}
	}

	return count
}

func (bst *BinarySearchTree) MaxSumAt() (int, int) {
	root := bst.rootNode
	if root == nil {
		return 0, 0
	}

	q := CreateQueue()
	MaxSum, MaxLevel := 0, 0
	CurSum, level := 0, 0

	q.Enqueue(root)
	q.Enqueue(nil)

	for !q.IsEmptyQueue() {
		temp := q.Dequeue()

		if temp == nil {
			if CurSum > MaxSum {
				MaxSum = CurSum
				MaxLevel = level
			}
			level++
			CurSum = 0
			if !q.IsEmptyQueue() {
				q.Enqueue(nil)
			}
		} else {
			CurSum += temp.key
			if temp.leftNode != nil {
				q.Enqueue(temp.leftNode)
			}

			if temp.rightNode != nil {
				q.Enqueue(temp.rightNode)
			}
		}
	}

	return MaxSum, MaxLevel
}

func (bst BinarySearchTree)PrintPathsToLeafNode() {

	root := bst.rootNode
	if root == nil {
		return
	}
	path := make([]int, 0)
	pathLen := 0

	path = append(path, root.key)
	pathLen++

	if root.leftNode == nil && root.rightNode == nil {
		fmt.Println(path, pathLen)
	} else {
		printPathsRecur(root.leftNode, path, pathLen)
		printPathsRecur(root.rightNode, path, pathLen)
	}
}

func printPathsRecur(root *TreeNode, path []int, pathLen int) {
	if root == nil {
		return
	}

	path = append(path, root.key)
	pathLen++
	

	if root.leftNode == nil && root.rightNode == nil {
		fmt.Println(path, pathLen)
	} else {
		printPathsRecur(root.leftNode, path, pathLen)
		printPathsRecur(root.rightNode, path, pathLen)
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
	bst.InsertElement(1)
	//bst.PreOrderTraversal()
	//bst.InOrderTraversal()
	//bst.PostOrderTraversal()
	//bst.LevelOrderTraversal()

	//fmt.Println (bst.FindMax())

	//fmt.Println(bst.SearchAnElement(12))
	//fmt.Println(bst.SearchAnElement(123))
	//fmt.Println(bst.SizeofBinaryTree())

	//bst.PrintLevelOrderinReverse()

	//fmt.Println(bst.LevelsofBinaryTree())

	//fmt.Println(bst.lastNodeInBT().key)

	//fmt.Println(bst.NumberOfFullNodes())
	//fmt.Println(bst.NumberOfHalfNodes())

	//fmt.Println(bst.MaxSumAt())

	bst.PrintPathsToLeafNode()
}
