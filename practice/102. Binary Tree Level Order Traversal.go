package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	if root == nil {
		return ans
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		item := []int{}
		size := len(queue)
		for i := 0; i < size; i++ {
			first := queue[0]
			item = append(item, first.Val)
			if first.Left != nil {
				queue = append(queue, first.Left)
			}
			if first.Right != nil {
				queue = append(queue, first.Right)
			}
			queue = queue[1:]
		}
		ans = append(ans, item)
	}

	return ans
}
