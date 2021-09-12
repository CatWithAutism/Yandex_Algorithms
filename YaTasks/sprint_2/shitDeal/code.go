package main

/*
Comment it before submitting
type ListNode struct {
    data   string
    next *ListNode
}
*/

func Solution(head *ListNode, idx int) *ListNode {
	if idx == 0 {
		head = head.next
		return head
	}

	prevNode := head
	nextNode := head
	for i := 0; i < idx; i++ {
		prevNode = nextNode
		nextNode = nextNode.next
	}

	if nextNode != nil && nextNode.next != nil {
		prevNode.next = nextNode.next
	} else {
		prevNode.next = nil
	}

	return head
}
