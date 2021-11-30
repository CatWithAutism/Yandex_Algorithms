package main

import "math"

func main() {
	test()
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func Solution(root *Node) bool {
	if root == nil {
		return true
	}

	leftHeight := getDepth(root.left)
	rightHeight := getDepth(root.right)
	return math.Abs(float64(leftHeight-rightHeight)) <= 1 && Solution(root.left) && Solution(root.right)
}

func getDepth(root *Node) int {
	if root == nil {
		return 0
	}

	return 1 + max(getDepth(root.left), getDepth(root.right))
}

func max(first, second int) int {
	if first > second {
		return first
	}

	return second
}

func test() {
	node1 := Node{1, nil, nil}
	node2 := Node{-5, nil, nil}
	node3 := Node{3, &node1, &node2}
	node4 := Node{10, nil, nil}
	node5 := Node{2, &node3, &node4}
	if !Solution(&node5) {
		panic("WA")
	}
}
