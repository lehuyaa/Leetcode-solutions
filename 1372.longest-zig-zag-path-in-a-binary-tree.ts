/*
 * @lc app=leetcode id=1372 lang=typescript
 *
 * [1372] Longest ZigZag Path in a Binary Tree
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     val: number
 *     left: TreeNode | null
 *     right: TreeNode | null
 *     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *     }
 * }
 */

function longestZigZag(root: TreeNode | null): number {
  let maxLength = 0;
  function dfs(
    root: TreeNode | null,
    direction: "L" | "R",
    currentLength: number
  ) {
    if (!root) {
      return;
    }
    maxLength = Math.max(maxLength, currentLength);
    dfs(root.left, "L", direction === "L" ? 1 : currentLength + 1);
    dfs(root.right, "R", direction === "R" ? 1 : currentLength + 1);
  }

  dfs(root, "L", 0);
  dfs(root, "R", 0);
  return maxLength;
}
// @lc code=end
