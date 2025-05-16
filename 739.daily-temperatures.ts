/*
 * @lc app=leetcode id=739 lang=typescript
 *
 * [739] Daily Temperatures
 */

// @lc code=start
function dailyTemperatures(temperatures: number[]): number[] {
  const n = temperatures.length;
  const wait = Array(n).fill(0);
  const stack: number[] = [];

  for (let i = n - 1; i >= 0; i--) {
    while (
      stack.length > 0 &&
      temperatures[stack[stack.length - 1]] <= temperatures[i]
    ) {
      stack.pop();
    }
    if (stack.length === 0) {
      wait[i] = 0;
    } else {
      wait[i] = stack[stack.length - 1] - i;
    }
    stack.push(i);
  }

  return wait;
}
// @lc code=end
