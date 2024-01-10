package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func amountOfTime(root *TreeNode, start int) int {
	graph := make(map[int][]int)
	buildGraph(root, graph)
	return bfs(graph, start)
}

func buildGraph(node *TreeNode, graph map[int][]int) {
	if node == nil {
		return
	}
	if node.Left != nil {
		graph[node.Val] = append(graph[node.Val], node.Left.Val)
		graph[node.Left.Val] = append(graph[node.Left.Val], node.Val)
	}
	if node.Right != nil {
		graph[node.Val] = append(graph[node.Val], node.Right.Val)
		graph[node.Right.Val] = append(graph[node.Right.Val], node.Val)
	}
	buildGraph(node.Left, graph)
	buildGraph(node.Right, graph)
}

func bfs(graph map[int][]int, start int) int {
	minutes := 0
	queue := []int{}
	queue = append(queue, start)
	visited := make(map[int]bool)
	visited[start] = true
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			element := queue[0]
			queue = queue[1:]
			for j := 0; j < len(graph[element]); j++ {
				if visited[graph[element][j]] {
					continue
				}
				queue = append(queue, graph[element][j])
				visited[graph[element][j]] = true
			}
		}
		minutes++
	}

	return minutes - 1
}
