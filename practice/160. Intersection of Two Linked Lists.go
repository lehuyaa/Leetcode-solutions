package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	visited := make(map[*ListNode]bool)
	for headA != nil {
		visited[headA] = true
		headA = headA.Next
	}
	for headB != nil {
		if visited[headB] {
			return headB
		}
		headB = headB.Next
	}

	return nil
}
