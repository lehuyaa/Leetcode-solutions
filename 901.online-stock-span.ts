/*
 * @lc app=leetcode id=901 lang=typescript
 *
 * [901] Online Stock Span
 */

// @lc code=start
class StockSpanner {
  prices: number[];
  stack: number[];
  size: number;
  constructor() {
    this.prices = [];
    this.stack = [];
    this.size = 0;
  }

  next(price: number): number {
    const currentIndex = this.size;
    this.prices.push(price);
    this.size++;
    while (
      this.stack.length > 0 &&
      this.prices[this.stack[this.stack.length - 1]] <= price
    ) {
      this.stack.pop();
    }

    const oldIndex =
      this.stack.length === 0 ? -1 : this.stack[this.stack.length - 1];
    this.stack.push(currentIndex);
    return currentIndex - oldIndex;
  }
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * var obj = new StockSpanner()
 * var param_1 = obj.next(price)
 */
// @lc code=end
