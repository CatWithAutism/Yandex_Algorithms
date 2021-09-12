package main

/*
Comment it before submitting
type ListNode struct {
    data   string
    next *ListNode
}
*/

func Solution(head *ListNode, elem string) int {
	i := 0
	for head != nil {
		if head.data == elem {
			return i
		}

		head = head.next
		i++
	}
	return -1
}
