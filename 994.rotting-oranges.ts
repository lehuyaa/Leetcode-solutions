/*
 * @lc app=leetcode id=994 lang=typescript
 *
 * [994] Rotting Oranges
 */

// @lc code=start
interface RottenOrange {
  x: number;
  y: number;
  minute: number;
}
const NEXT_TO = [
  [-1, 0],
  [1, 0],
  [0, -1],
  [0, 1],
];
function orangesRotting(grid: number[][]): number {
  const row = grid.length;
  const col = grid[0].length;
  const times = Array(row)
    .fill(null)
    .map(() => Array(col).fill(0));
  const queue: RottenOrange[] = [];
  for (let i = 0; i < row; i++) {
    for (let j = 0; j < col; j++) {
      if (grid[i][j] === 2) {
        queue.push({ x: i, y: j, minute: 0 });
      }
    }
  }

  while (queue.length > 0) {
    const { x: currentX, y: currentY, minute: currentMinute } = queue.shift()!;

    for (const [nextX, nextY] of NEXT_TO) {
      const newX = currentX + nextX;
      const newY = currentY + nextY;
      if (
        newX < 0 ||
        newX >= row ||
        newY < 0 ||
        newY >= col ||
        grid[newX][newY] === 0 ||
        grid[newX][newY] === 2
      ) {
        continue;
      }
      queue.push({ x: newX, y: newY, minute: currentMinute + 1 });
      grid[newX][newY] = 2;
      times[newX][newY] = currentMinute + 1;
    }
  }
  let maxTime = 0;
  for (let i = 0; i < row; i++) {
    for (let j = 0; j < col; j++) {
      if (grid[i][j] === 1) {
        return -1;
      }
      if (times[i][j] > maxTime) {
        maxTime = times[i][j];
      }
    }
  }

  return maxTime;
}
// @lc code=end
