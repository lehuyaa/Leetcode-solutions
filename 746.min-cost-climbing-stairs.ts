/*
 * @lc app=leetcode id=746 lang=typescript
 *
 * [746] Min Cost Climbing Stairs
 */

// @lc code=start
function minCostClimbingStairs(cost: number[]): number {
  const n = cost.length;
  const dp = Array(n + 1).fill(0);

  for (let i = 2; i <= n; i++) {
    dp[i] = Math.min(dp[i - 1] + cost[i - 1], dp[i - 2] + cost[i - 2]);
  }

  return dp[n];
}
// @lc code=end
