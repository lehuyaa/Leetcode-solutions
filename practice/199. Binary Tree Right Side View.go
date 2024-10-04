package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	ans := []int{}
	if root == nil {
		return ans
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		newQueue := []*TreeNode{}
		ans = append(ans, queue[len(queue)-1].Val)
		for i := 0; i < len(queue); i++ {
			if queue[i].Left != nil {
				newQueue = append(newQueue, queue[i].Left)
			}
			if queue[i].Right != nil {
				newQueue = append(newQueue, queue[i].Right)
			}
		}
		queue = newQueue
	}
	return ans
}
