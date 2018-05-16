package main

import "fmt"

// Node : asdas
type Node struct {
	Key int
}

func dijkstra(src int, nodes map[int]*Node, matrix [][]int) (dist, prev []int) {
	unvisited := make(map[int]*Node)
	for k, v := range nodes {
		unvisited[k] = v
	}
	dist = make([]int, len(nodes))
	prev = make([]int, len(nodes))
	for i := 0; i < len(nodes); i++ {
		dist[i] = 1000000
		prev[i] = -1
	}
	dist[src] = 0
	for len(unvisited) > 0 {
		min := -1
		for i := range unvisited {
			if min == -1 || dist[i] < dist[min] {
				min = i
			}
		}
		unv := unvisited[min]
		delete(unvisited, min)

		for i := range matrix[unv.Key] {
			alt := dist[unv.Key] + matrix[unv.Key][i]
			if alt < dist[i] {
				dist[i] = alt
				prev[i] = unv.Key
			}
			fmt.Println(alt)
		}
	}
	return
}

func main() {
	nodes := map[int]*Node{
		0: {Key: 0},
		1: {Key: 1},
		2: {Key: 2},
		3: {Key: 3},
		4: {Key: 4},
	}

	matrix := [][]int{
		{0, 1, 5, 2, 10},
		{1, 0, 3, 1, 5},
		{5, 3, 0, 6, 6},
		{2, 1, 6, 0, 5},
		{10, 5, 6, 5, 0},
	}

	dist, prev := dijkstra(0, nodes, matrix)
	fmt.Println("Distances from node 0:")
	fmt.Println(dist)
	fmt.Println("Positions previous to node 0:")
	fmt.Println(prev)
}
