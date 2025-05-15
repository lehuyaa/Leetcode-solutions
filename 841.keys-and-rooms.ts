/*
 * @lc app=leetcode id=841 lang=typescript
 *
 * [841] Keys and Rooms
 */

// @lc code=start
function canVisitAllRooms(rooms: number[][]): boolean {
  const visited = Array(rooms.length).fill(false);

  function dfs(room: number) {
    visited[room] = true;

    const keys = rooms[room];
    for (const key of keys) {
      if (!visited[key]) {
        dfs(key);
      }
    }
  }
  dfs(0);
  return visited.every((item) => !!item);
}
// @lc code=end
