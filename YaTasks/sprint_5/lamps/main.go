// Lamps Лампочки
package main

func main() {
	test()
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func Solution(root *Node) int {
	if root.left != nil && root.right != nil {
		return max(root.value, max(Solution(root.left), Solution(root.right)))
	} else if root.left != nil {
		return max(root.value, Solution(root.left))
	} else if root.right != nil {
		return max(root.value, Solution(root.right))
	}

	return root.value
}

func max(first, second int) int {
	if first >= second {
		return first
	}

	return second
}

func test() {
	node1 := Node{1, nil, nil}
	node2 := Node{-5, nil, nil}
	node3 := Node{3, &node1, &node2}
	node4 := Node{2, &node3, nil}
	if Solution(&node4) != 3 {
		panic("WA")
	}
}
