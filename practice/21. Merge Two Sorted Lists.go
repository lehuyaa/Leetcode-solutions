package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	curr := head
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}

	for list1 != nil {
		curr.Next = &ListNode{list1.Val, nil}
		list1 = list1.Next
		curr = curr.Next
	}

	for list2 != nil {
		curr.Next = &ListNode{list2.Val, nil}
		list2 = list2.Next
		curr = curr.Next
	}

	return head.Next
}
