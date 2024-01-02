package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return createNode(nums, 0, len(nums)-1)
}

func createNode(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := left + (right-left)/2
	nodeLeft := createNode(nums, left, mid-1)
	nodeRight := createNode(nums, mid+1, right)
	return &TreeNode{nums[mid], nodeLeft, nodeRight}

}
