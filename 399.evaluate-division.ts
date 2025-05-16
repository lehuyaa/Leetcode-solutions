/*
 * @lc app=leetcode id=399 lang=typescript
 *
 * [399] Evaluate Division
 */

// @lc code=start
function calcEquation(
  equations: string[][],
  values: number[],
  queries: string[][]
): number[] {
  const neighbors = new Map<string, { next: string; point: number }[]>();
  for (let i = 0; i < equations.length; i++) {
    const [a, b] = equations[i];
    neighbors.set(
      a,
      neighbors.has(a)
        ? [...neighbors.get(a)!, { next: b, point: values[i] }]
        : [{ next: b, point: values[i] }]
    );
    if (values[i] !== 0) {
      neighbors.set(
        b,
        neighbors.has(b)
          ? [...neighbors.get(b)!, { next: a, point: 1 / values[i] }]
          : [{ next: a, point: 1 / values[i] }]
      );
    }
  }
  function count(from: string, to: string): number {
    if (!neighbors.has(from) || !neighbors.has(to)) {
      return -1;
    }
    const visited = new Map<string, boolean>();
    const queue: { title: string; value: number }[] = [];
    queue.push({ title: from, value: 1 });
    while (queue.length > 0) {
      const { title: currentTitle, value: currentValue } = queue.shift()!;
      if (currentTitle === to) {
        return currentValue;
      }
      visited.set(currentTitle, true);
      for (const { next, point } of neighbors.get(currentTitle) || []) {
        if (!visited.has(next)) {
          queue.push({ title: next, value: currentValue * point });
        }
      }
    }

    return -1;
  }
  const ans = Array(queries.length).fill(0);
  for (let i = 0; i < ans.length; i++) {
    ans[i] = count(queries[i][0], queries[i][1]);
  }
  return ans;
}
// @lc code=end
