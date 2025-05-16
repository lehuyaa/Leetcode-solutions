/*
 * @lc app=leetcode id=72 lang=typescript
 *
 * [72] Edit Distance
 */

// @lc code=start
function minDistance(word1: string, word2: string): number {
  //  "" r o s
  //"" 0 1 2 3
  // h 1 1 2 3
  // o 2 2 1 2
  // r 3 2 2 2
  // s 4 3 3 2
  // e 5 4 4 3

  const len1 = word1.length;
  const len2 = word2.length;
  const dp = Array(len2 + 1)
    .fill(null)
    .map(() => Array(len1 + 1).fill(0));

  for (let i = 0; i < len1; i++) {
    dp[0][i] = i;
  }
  for (let j = 0; j < len2; j++) {
    dp[j][0] = j;
  }
  for (let i = 1; i <= len2; i++) {
    for (let j = 1; j <= len1; j++) {
      if (word1[j - 1] === word2[i - 1]) {
        dp[i][j] = dp[i - 1][j - 1];
      } else {
        dp[i][j] = 1 + Math.min(dp[i - 1][j - 1], dp[i][j - 1], dp[i - 1][j]);
      }
    }
  }

  return dp[len2][len1];
}
// @lc code=end
