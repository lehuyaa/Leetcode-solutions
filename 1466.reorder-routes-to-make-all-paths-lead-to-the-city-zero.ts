/*
 * @lc app=leetcode id=1466 lang=typescript
 *
 * [1466] Reorder Routes to Make All Paths Lead to the City Zero
 */

// @lc code=start
function minReorder(n: number, connections: number[][]): number {
  // 0 -> [1,T]
  //      [4,F]
  // 1 -> [3,T]
  //      [0,F]
  // 2 -> [3,T]
  // 3 -> [1,F]
  //      [2,F]
  // 4 -> [0,T]
  // 5 -> [4,F]
  const neighbors = new Map<number, { next: number; isForward: boolean }[]>();
  for (const [a, b] of connections) {
    if (!neighbors.has(a)) neighbors.set(a, []);
    if (!neighbors.has(b)) neighbors.set(b, []);

    neighbors.get(a)?.push({ next: b, isForward: true });
    neighbors.get(b)?.push({ next: a, isForward: false });
  }
  const visited = Array(n).fill(false);
  const queue: number[] = [];
  let needChange = 0;
  queue.push(0);
  while (queue.length > 0) {
    const top = queue.shift()!;
    const neighbor = neighbors.get(top) || [];
    visited[top] = true;
    for (const { next, isForward } of neighbor) {
      if (!visited[next]) {
        if (isForward) {
          needChange++;
        }
        queue.push(next);
      }
    }
  }

  return needChange;
}
// @lc code=end
