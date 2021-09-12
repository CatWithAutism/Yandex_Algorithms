package main

import (
	"bufio"
	"os"
)

/*
type ListNode struct {
	data   string
	next *ListNode
}
*/

func Solution(head *ListNode) {
	for head != nil {
		writeData(head.data)
		head = head.next
	}
}

func writeData(data string) {
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString(data + "\n")
	writer.Flush()
}
