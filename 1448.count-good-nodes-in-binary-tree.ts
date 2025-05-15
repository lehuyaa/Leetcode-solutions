/*
 * @lc app=leetcode id=1448 lang=typescript
 *
 * [1448] Count Good Nodes in Binary Tree
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

function goodNodes(root: TreeNode | null): number {
  if (!root) {
    return 0;
  }
  let countGood = 0;
  function findNode(root: TreeNode | null, max: number) {
    if (!root) {
      return;
    }
    let maxValue = max;
    if (root.val >= max) {
      countGood++;
      maxValue = root.val;
    }
    findNode(root.left, maxValue);
    findNode(root.right, maxValue);
  }
  findNode(root, root.val);
  return countGood;
}
// @lc code=end
