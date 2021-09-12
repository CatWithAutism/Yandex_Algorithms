package main

import (
	"bufio"
	"os"
)

/*
Comment it before submitting
type ListNode struct {
    data     string
    next  *ListNode
    prev  *ListNode
}

*/

func Solution(head *ListNode) *ListNode {
	for {
		if head.next == nil {
			break
		}
		head = head.next
	}

	currentNode := head
	var temp *ListNode
	for {
		temp = currentNode.prev
		currentNode.prev = currentNode.next
		currentNode.next = temp

		if currentNode.next == nil {
			break
		}

		currentNode = currentNode.next
	}
	return head
}

func writeData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data + "\n")
	writer.Flush()
}
