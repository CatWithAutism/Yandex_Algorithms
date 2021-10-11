package main

func main() {
	test()
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func isBTS(root *Node, rootVal int, compareFunc func(int, int) bool) bool {
	if root == nil {
		return true
	}

	stack := NewStack()
	stack.Push(root)

	for currentElement := stack.Pop(); currentElement != nil; currentElement = stack.Pop() {
		if currentElement.left != nil {
			if currentElement.left.value >= currentElement.value {
				return false
			}

			stack.Push(currentElement.left)
		}

		if currentElement.right != nil {
			if currentElement.right.value <= currentElement.value {
				return false
			}

			stack.Push(currentElement.right)
		}

		if !compareFunc(rootVal, currentElement.value) {
			return false
		}
	}

	return true
}

func Solution(root *Node) bool {
	return isBTS(root.left, root.value, func(i int, i2 int) bool {
		return i > i2
	}) && isBTS(root.right, root.value, func(i int, i2 int) bool {
		return i < i2
	})
}

func test() {
	node1 := Node{1, nil, nil}
	node2 := Node{4, nil, nil}
	node3 := Node{3, &node1, &node2}
	node4 := Node{8, nil, nil}
	node5 := Node{5, &node3, &node4}
	if !Solution(&node5) {
		panic("WA")
	}
	node2.value = 5
	if Solution(&node5) {
		panic("WA")
	}
}

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value *Node
		prev  *node
	}
)

func NewStack() *Stack {
	return &Stack{nil, 0}
}

func (stack *Stack) Len() int {
	return stack.length
}

func (stack *Stack) Pop() *Node {
	if stack.length == 0 {
		return nil
	}

	n := stack.top
	stack.top = n.prev
	stack.length--
	return n.value
}

func (stack *Stack) Push(value *Node) {
	n := &node{value, stack.top}
	stack.top = n
	stack.length++
}
