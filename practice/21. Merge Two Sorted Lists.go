package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	cur := head
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = &ListNode{
				Val:  list1.Val,
				Next: nil,
			}
			cur = cur.Next
			list1 = list1.Next
		} else {
			cur.Next = &ListNode{
				Val:  list1.Val,
				Next: nil,
			}
			cur = cur.Next
			list2 = list1.Next
		}
	}

	for list1 != nil {
		cur.Next = &ListNode{
			Val:  list1.Val,
			Next: nil,
		}
		cur = cur.Next
		list1 = list1.Next
	}

	for list2 != nil {
		cur.Next = &ListNode{
			Val:  list2.Val,
			Next: nil,
		}
		cur = cur.Next
		list2 = list2.Next
	}

	return head.Next
}
