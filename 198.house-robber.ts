/*
 * @lc app=leetcode id=198 lang=typescript
 *
 * [198] House Robber
 */

// @lc code=start
function rob(nums: number[]): number {
  const n = nums.length;
  const dp = Array(n)
    .fill(null)
    .map(() => Array(2).fill(0));
  dp[0][0] = 0;
  dp[0][1] = nums[0];
  for (let i = 1; i < n; i++) {
    dp[i][0] = dp[i - 1][1];
    dp[i][1] = Math.max(dp[i - 1][1], dp[i - 1][0] + nums[i]);
  }

  return Math.max(dp[n - 1][0], dp[n - 1][1]);
}
// @lc code=end
