/*
 * @lc app=leetcode id=1926 lang=typescript
 *
 * [1926] Nearest Exit from Entrance in Maze
 */

// @lc code=start
interface Place {
  x: number;
  y: number;
  d: number;
}
const NEXT_TO = [
  [-1, 0],
  [1, 0],
  [0, -1],
  [0, 1],
];
function nearestExit(maze: string[][], entrance: number[]): number {
  const row = maze.length;
  const col = maze[0].length;

  const queue: Place[] = [];
  queue.push({ x: entrance[0], y: entrance[1], d: 0 });
  maze[entrance[0]][entrance[1]] = "+";

  while (queue.length > 0) {
    const { x: currentX, y: currentY, d: currentDistance } = queue.shift()!;
    for (const [nextX, nextY] of NEXT_TO) {
      const newX = currentX + nextX;
      const newY = currentY + nextY;
      if (
        newX < 0 ||
        newX >= row ||
        newY < 0 ||
        newY >= col ||
        maze[newX][newY] === "+"
      ) {
        continue;
      }
      const isExit =
        newX === 0 || newY === 0 || newX === row - 1 || newY === col - 1;
      if (isExit) {
        return currentDistance + 1;
      }
      queue.push({ x: newX, y: newY, d: currentDistance + 1 });
      maze[newX][newY] = "+";
    }
  }

  return -1;
}
// [["+","+","+"],[".",".","."],["+","+","+"]]\n[1,0]
// @lc code=end
