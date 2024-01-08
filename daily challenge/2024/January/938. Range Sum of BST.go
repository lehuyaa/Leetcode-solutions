package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBST(node *TreeNode, l int, h int) int {
	if node == nil {
		return 0
	}
	count := 0
	if node.Val >= l && node.Val <= h {
		count += node.Val
	}
	if node.Val < l {
		count += findBST(node.Right, l, h)
	} else if node.Val > h {
		count += findBST(node.Left, l, h)
	} else {
		count += findBST(node.Right, l, h)
		count += findBST(node.Left, l, h)
	}
	return count
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	return findBST(root, low, high)
}
