/*
	Сложность: O(h)
	Получаем на вход ключ и вершину.
	Т.к. это бинарное дерево мы знаем с какой стороны искать узел с ключем
	Находим ключ и удаляем узел в соотвествии с условиями

	Если узел последний и у него нету потомков, то удаляем забываем.
	Если узел содержит в себе только левого или правого потомка, то удаляем, а узел потомка пробрасываем выше.
	Если узел содержит оба потомка, то начинаем танцы с бубном:
		Берем максимальное правое значение левого потомка, либо минимальное левое правого потомка.
		Это у нас и будет нашей заменяемой нодой

	Посылка - https://contest.yandex.ru/contest/24810/run-report/56564137/
*/

package main

type Node struct {
	value int
	left  *Node
	right *Node
}

func remove(node *Node, key int) *Node {
	if node == nil {
		return node
	}

	if node.value > key {
		node.left = remove(node.left, key)
	} else if node.value < key {
		node.right = remove(node.right, key)
	} else if node.value == key {
		if node.left == nil && node.right == nil {
			return nil
		} else if node.left == nil && node.right != nil {
			return node.right
		} else if node.left != nil && node.right == nil {
			return node.left
		} else {
			rightHandLeftMinNode := getLeftMin(node.right)
			node.value = rightHandLeftMinNode.value
			node.right = remove(node.right, rightHandLeftMinNode.value)
		}
	}

	return node
}

func getLeftMin(root *Node) *Node {
	currentNode := root
	for nil != currentNode && currentNode.left != nil {
		currentNode = currentNode.left
	}
	return currentNode
}

func test() {
	node1 := Node{2, nil, nil}
	node2 := Node{3, &node1, nil}
	node3 := Node{1, nil, &node2}
	node4 := Node{6, nil, nil}
	node5 := Node{8, &node4, nil}
	node6 := Node{10, &node5, nil}
	node7 := Node{5, &node3, &node6}
	newHead := remove(&node7, 10)
	if newHead.value != 5 {
		panic("WA")
	}
	if newHead.right != &node5 {
		panic("WA")
	}
	if newHead.right.value != 8 {
		panic("WA")
	}
}

func main() {
	test()
}
