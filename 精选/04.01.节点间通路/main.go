package main

var visit []bool

func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	visit = make([]bool, len(graph))

	return dfs(graph, start, target)
}

func dfs(graph [][]int, start int, target int) bool {

	for index := range graph {
		if !visit[index] {
			if graph[index][0] == start && graph[index][1] == target {
				return true
			}

			visit[index] = true

			if graph[index][1] == target && dfs(graph, start, graph[index][0]) {
				return true
			}

			visit[index] = false
		}
	}

	return false
}
