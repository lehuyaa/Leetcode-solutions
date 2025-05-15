/*
 * @lc app=leetcode id=1161 lang=typescript
 *
 * [1161] Maximum Level Sum of a Binary Tree
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

function maxLevelSum(root: TreeNode | null): number {
  if (!root) {
    return 0;
  }

  const level: number[][] = [];
  const queue: TreeNode[] = [];
  queue.push(root);

  while (queue.length > 0) {
    const size = queue.length;
    const currentLevel: number[] = [];
    for (let i = 0; i < size; i++) {
      const top = queue.shift();
      currentLevel.push(top.val);
      if (top.left) {
        queue.push(top.left);
      }
      if (top.right) {
        queue.push(top.right);
      }
    }
    level.push(currentLevel);
  }
  let max = -Infinity;
  let ans = 0;
  for (let i = 0; i < level.length; i++) {
    const sum = level[i].reduce((prev, cur) => prev + cur, 0);
    if (sum > max) {
      max = sum;
      ans = i;
    }
  }

  return ans + 1;
}
// @lc code=end
