package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans []int

func postorderTraversal(root *TreeNode) []int {
	ans = []int{}
	if root == nil {
		return ans
	}
	postorder(root)
	return ans
}

func postorder(root *TreeNode) {
	if root.Left == nil && root.Right == nil {
		ans = append(ans, root.Val)
		return
	}
	if root.Left != nil {
		postorder(root.Left)
	}
	if root.Right != nil {
		postorder(root.Right)
	}

	ans = append(ans, root.Val)
}
