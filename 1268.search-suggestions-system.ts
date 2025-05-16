/*
 * @lc app=leetcode id=1268 lang=typescript
 *
 * [1268] Search Suggestions System
 */

// @lc code=start
function suggestedProducts(products: string[], searchWord: string): string[][] {
  // ["mobile","mouse","moneypot","monitor","mousepad"]
  // "" - []
  // m - ["mobile","mouse","moneypot","monitor","mousepad"]
  // o - ["mobile", "mouse", "moneypot","monitor","mousepad"]
  // b - [mobile], u - [mouse, mousepad], n - [moneypot, monitor]
  //
  products.sort();
  let maxLen = 0;
  for (const product of products) {
    if (product.length > maxLen) {
      maxLen = product.length;
    }
  }
  const mapping = Array(maxLen)
    .fill(null)
    .map(() => new Map<string, string[]>());
  for (let i = 0; i < maxLen; i++) {
    for (const product of products) {
      if (i >= product.length) {
        continue;
      }
    }
  }
}
// @lc code=end
