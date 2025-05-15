/*
 * @lc app=leetcode id=872 lang=typescript
 *
 * [872] Leaf-Similar Trees
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

function leafSimilar(root1: TreeNode | null, root2: TreeNode | null): boolean {
  const leaf1: number[] = [];
  const leaf2: number[] = [];

  function dfs(root: TreeNode | null, pool: number[]) {
    if (root === null) {
      return;
    }
    if (!root.left && !root.right) {
      pool.push(root.val);
    }

    dfs(root.left, pool);
    dfs(root.right, pool);
  }

  dfs(root1, leaf1);
  dfs(root2, leaf2);
  if (leaf1.length !== leaf2.length) {
    return false;
  }

  for (let i = 0; i < leaf1.length; i++) {
    if (leaf1[i] !== leaf2[i]) {
      return false;
    }
  }

  return true;
}
// @lc code=end
