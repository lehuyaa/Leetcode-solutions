/*
 * @lc app=leetcode id=437 lang=typescript
 *
 * [437] Path Sum III
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

function pathSum(root: TreeNode | null, targetSum: number): number {
  if (!root) return 0;
  function countPath(root: TreeNode | null, targetSum: number): number {
    if (!root) return 0;
    let count = 0;

    if (root.val === targetSum) count++;

    count += countPath(root.left, targetSum - root.val);
    count += countPath(root.right, targetSum - root.val);

    return count;
  }

  return (
    countPath(root, targetSum) +
    pathSum(root.left, targetSum) +
    pathSum(root.right, targetSum)
  );
}
// @lc code=end
