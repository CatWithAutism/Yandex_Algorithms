package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	test()
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func printRange(root *Node, left int, right int) {
	if root == nil {
		return
	}

	//обходим слева пока выполняется условие
	if left <= root.value {
		printRange(root.left, left, right)
	}

	//на последнем элементе начнем печать
	//т.е. выйдет => min -> max
	if left <= root.value && right >= root.value {
		writeData(strconv.Itoa(root.value))
	}

	//обходим элементы больше
	if right >= root.value {
		printRange(root.right, left, right)
	}
}

func writeData(data string) {
	consoleWriter := bufio.NewWriter(os.Stdout)
	consoleWriter.WriteString(data + " ")
	consoleWriter.Flush()
}

func test() {
	node1 := Node{2, nil, nil}
	node2 := Node{1, nil, &node1}
	node3 := Node{8, nil, nil}
	node4 := Node{8, nil, &node3}
	node5 := Node{9, &node4, nil}
	node6 := Node{10, &node5, nil}
	node7 := Node{5, &node2, &node6}
	printRange(&node7, 2, 8)
	// expected output: 2 5 8 8
}

func test1() {
	node10 := Node{932, nil, nil}
	node9 := Node{912, nil, &node10}
	node8 := Node{822, nil, nil}
	node7 := Node{870, &node8, &node9}
	node6 := Node{701, nil, nil}
	node5 := Node{702, &node6, &node7}
	node4 := Node{266, nil, nil}
	node3 := Node{191, nil, &node4}
	node2 := Node{298, &node3, nil}
	node1 := Node{668, &node2, &node5}
	printRange(&node1, 73, 545)
	//191
	//266
	//298
}
