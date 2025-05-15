/*
 * @lc app=leetcode id=547 lang=typescript
 *
 * [547] Number of Provinces
 */

// @lc code=start
function findCircleNum(isConnected: number[][]): number {
  const numberOfCity = isConnected.length;
  const parents = Array(numberOfCity)
    .fill(null)
    .map((_, index) => index);

  function findParent(node: number): number {
    if (parents[node] === node) {
      return node;
    }

    return findParent(parents[node]);
  }

  for (let i = 0; i < numberOfCity; i++) {
    for (let j = 0; j < numberOfCity; j++) {
      if (isConnected[i][j] === 1) {
        const pI = findParent(i);
        const pJ = findParent(j);
        if (pI !== pJ) {
          parents[pI] = pJ;
        }
      }
    }
  }

  let count = 0;
  for (let i = 0; i < parents.length; i++) {
    if (parents[i] === i) {
      count++;
    }
  }
  return count;
}
// @lc code=end
