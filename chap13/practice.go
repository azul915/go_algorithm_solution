package chap13

import (
	helper "algorithm_solution/pkg"
	"fmt"
)

type Graph [][]int
type Search struct{}

func Code1() {

	g := &Graph{{1, 2, 4}, {0, 3, 8}, {0, 5}, {1, 7, 8}, {0, 8}, {2, 6, 8}, {5, 7}, {3, 6}, {1, 3, 4, 5}}
	s := &Search{}
	s.Bfs(g, 0)
	fmt.Println("")
	s.Dfs(g, 0)
}

// bfs
func (sh *Search) Bfs(G *Graph, s int) {
	N := len(*G)

	// seen
	seen := make([]bool, N)
	// todo
	todo := helper.NewQueue[int]()

	seen[s] = true
	fmt.Printf("visited: %v\n", s)
	todo.Enqueue(s)

	for !todo.Empty() {
		v := todo.Dequeue()
		for _, v := range (*G)[v] {
			if seen[v] {
				continue
			}
			seen[v] = true
			fmt.Printf("visited: %v\n", v)
			todo.Enqueue(v)
		}
	}
}

// dfs
func (sh *Search) Dfs(G *Graph, s int) {
	N := len(*G)

	// seen
	seen := make([]bool, N)
	// todo
	todo := helper.NewStack[int]()

	seen[s] = true
	fmt.Printf("visited: %v\n", s)
	todo.Push(s)

	for !todo.Empty() {
		v := todo.Pop()
		for _, v := range (*G)[v] {
			if seen[v] {
				continue
			}
			seen[v] = true
			fmt.Printf("visited: %v\n", v)
			todo.Push(v)
		}
	}
}
