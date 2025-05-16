/*
 * @lc app=leetcode id=790 lang=typescript
 *
 * [790] Domino and Tromino Tiling
 */

// @lc code=start
const MOD = 1e9 + 7;
function numTilings(n: number): number {
  if (n === 1) {
    return 1;
  }
  if (n === 2) {
    return 2;
  }
  const sum = Array(n + 1).fill(0);
  const dp = Array(n + 1).fill(0);
  dp[0] = 1;
  sum[0] = 1;
  dp[1] = 1;
  sum[1] = 2;
  dp[2] = 2;
  sum[2] = 4;
  for (let i = 3; i <= n; i++) {
    dp[i] = (dp[i - 1] + dp[i - 2] + 2 * sum[i - 3]) % MOD;
    sum[i] = sum[i - 1] + dp[i];
  }

  return dp[n];
}
// @lc code=end
