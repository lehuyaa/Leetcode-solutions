/*
 * @lc app=leetcode id=714 lang=typescript
 *
 * [714] Best Time to Buy and Sell Stock with Transaction Fee
 */

// @lc code=start
function maxProfit(prices: number[], fee: number): number {
  // [1,3,2,8,4,9], fee = 2
  // 0: hold
  // 1: profit
  // hold at i = max(hold[i-1], profit[i-1] - prices[i])
  // profit at i = max(profit[i-1], hold[i-1] + prices[i] - fee)
  // 0: [-1,0] 1
  // 1: [-1,0] 3
  // 2: [-1,0] 2
  // 3: [-1,5] 8
  // 4: [1,5] 4
  // 5: [1,8] 9

  const n = prices.length;
  const hold = Array(n).fill(0);
  const profit = Array(n).fill(0);
  hold[0] = -prices[0];
  for (let i = 1; i < n; i++) {
    hold[i] = Math.max(hold[i - 1], profit[i - 1] - prices[i]);
    profit[i] = Math.max(profit[i - 1], hold[i - 1] + prices[i] - fee);
  }
  return Math.max(hold[n - 1], profit[n - 1]);
}
// @lc code=end
