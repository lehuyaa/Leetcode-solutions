/*
 * @lc app=leetcode id=199 lang=typescript
 *
 * [199] Binary Tree Right Side View
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

function rightSideView(root: TreeNode | null): number[] {
  if (!root) {
    return [];
  }
  const ans: number[] = [];
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

  for (let i = 0; i < level.length; i++) {
    if (level[i].length > 0) {
      ans.push(level[i][level[i].length - 1]);
    }
  }

  return ans;
}
// @lc code=end
