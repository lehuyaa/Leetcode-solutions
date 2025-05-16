/*
 * @lc app=leetcode id=62 lang=typescript
 *
 * [62] Unique Paths
 */

// @lc code=start
const MOD = 2e9;
function uniquePaths(m: number, n: number): number {
  const dp = Array(m)
    .fill(null)
    .map(() => Array(n).fill(0));

  for (let i = 0; i < m; i++) {
    for (let j = 0; j < n; j++) {
      if (i === 0 && j === 0) {
        dp[i][j] = 1;
      } else {
        const top = i === 0 ? 0 : dp[i - 1][j];
        const left = j === 0 ? 0 : dp[i][j - 1];
        dp[i][j] = (top + left) % MOD;
      }
    }
  }

  return dp[m - 1][n - 1];
}
// @lc code=end
