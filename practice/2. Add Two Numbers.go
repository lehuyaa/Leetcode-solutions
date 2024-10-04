package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{
		Val:  0,
		Next: nil,
	}
	cur := head
	num := 0
	for l1 != nil && l2 != nil {
		number := num + l1.Val + l2.Val
		if number >= 10 {
			num = 1
			number -= 10
		} else {
			num = 0
		}
		cur.Next = &ListNode{
			Val:  number,
			Next: nil,
		}
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		number := num + l1.Val
		if number >= 10 {
			num = 1
			number -= 10
		} else {
			num = 0
		}
		cur.Next = &ListNode{
			Val:  number,
			Next: nil,
		}
		cur = cur.Next
		l1 = l1.Next
	}
	for l2 != nil {
		number := num + l2.Val
		if number >= 10 {
			num = 1
			number -= 10
		} else {
			num = 0
		}
		cur.Next = &ListNode{
			Val:  number,
			Next: nil,
		}
		cur = cur.Next
		l2 = l2.Next
	}

	if num == 1 {
		cur.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
		cur = cur.Next
	}

	return head.Next
}
