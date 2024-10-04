package main

var visited [][]bool

func exist(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])

	visited = make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] && dfs(i, j, board, word, 0) {
				return true
			}
		}
	}

	return false
}

func dfs(x int, y int, board [][]byte, word string, index int) bool {
	if index == len(word) {
		return true
	}

	if x < 0 || y < 0 || x >= len(board) || y >= len(board[0]) {
		return false
	}

	if visited[x][y] || board[x][y] != word[index] {
		return false
	}

	visited[x][y] = true

	if dfs(x-1, y, board, word, index+1) || dfs(x, y+1, board, word, index+1) || dfs(x+1, y, board, word, index+1) || dfs(x, y-1, board, word, index+1) {
		return true
	}

	visited[x][y] = false

	return false
}
