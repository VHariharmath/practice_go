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

type BinaryTree struct {
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

func (bt *BinaryTree) InsertElement(key int) {
	newNode := &TreeNode{key, nil, nil}
	if bt.rootNode == nil {
		bt.rootNode = newNode
		return
	} else {
		insertNode(bt.rootNode, newNode)
	}
}

func (bt *BinaryTree) PreOrderTraversal() {
	s := CreateStack()

	root := bt.rootNode

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

func (bt *BinaryTree) InOrderTraversal() {
	s := CreateStack()
	root := bt.rootNode

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

func (bt *BinaryTree) PostOrderTraversal() {
	s := CreateStack()
	root := bt.rootNode

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

func (bt *BinaryTree) LevelOrderTraversal() {
	q := CreateQueue()

	root := bt.rootNode
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

func (bt *BinaryTree) FindMax() int {
	q := CreateQueue()

	root := bt.rootNode
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

func (bt *BinaryTree) SearchAnElement(key int) bool {
	q := CreateQueue()

	root := bt.rootNode
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
func (bt *BinaryTree) InsertElement(key int) {
	newNode := &TreeNode{}
	newNode.key = key

	root := bt.rootNode
	if root == nil {
		bt.rootNode = newNode
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

func (bt *BinaryTree) SizeofBinaryTree() int {

	root := bt.rootNode
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

func (bt *BinaryTree) PrintLevelOrderinReverse() {

	root := bt.rootNode
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

func (bt *BinaryTree) LevelsofBinaryTree() int {

	root := bt.rootNode
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

func (bt *BinaryTree) lastNodeInBT() *TreeNode {
	q := CreateQueue()
	temp := &TreeNode{}
	root := bt.rootNode

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

func (bt *BinaryTree) NumberOfFullNodes() int {
	root := bt.rootNode
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

func (bt *BinaryTree) NumberOfHalfNodes() int {
	root := bt.rootNode
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

func (bt *BinaryTree) MaxSumAt() (int, int) {
	root := bt.rootNode
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

func (bt BinaryTree)PrintPathsToLeafNode() {

	root := bt.rootNode
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
	bt := &BinaryTree{}
	bt.InsertElement(10)
	bt.InsertElement(9)
	bt.InsertElement(11)
	bt.InsertElement(8)
	bt.InsertElement(15)
	bt.InsertElement(12)
	bt.InsertElement(5)
	bt.InsertElement(7)
	bt.InsertElement(6)
	bt.InsertElement(1)
	//bt.PreOrderTraversal()
	bt.InOrderTraversal()
	//bt.PostOrderTraversal()
	//bt.LevelOrderTraversal()

	//fmt.Println (bt.FindMax())

	//fmt.Println(bt.SearchAnElement(12))
	//fmt.Println(bt.SearchAnElement(123))
	//fmt.Println(bt.SizeofBinaryTree())

	//bt.PrintLevelOrderinReverse()

	//fmt.Println(bt.LevelsofBinaryTree())

	//fmt.Println(bt.lastNodeInBT().key)

	//fmt.Println(bt.NumberOfFullNodes())
	//fmt.Println(bt.NumberOfHalfNodes())

	//fmt.Println(bt.MaxSumAt())

	//bt.PrintPathsToLeafNode()
}
