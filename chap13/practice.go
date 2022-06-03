package chap13

import (
	"fmt"

	"github.com/azul915/go_algorithm_solution/pkg/helper"
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

func Code2() {
	seen := make([]bool, 8)
	G := &Graph{{5}, {6, 3}, {5, 7}, {0, 7}, {1, 6}, {}, {7}, {0}}

	for v := 0; v < len(*G); v++ {
		if seen[v] {
			continue
		}
		dfs(G, v, seen)
	}
}

func dfs(G *Graph, v int, seen []bool) {
	seen[v] = true
	fmt.Printf("visited: %v\n", v)
	for _, next := range (*G)[v] {
		if seen[next] {
			continue
		}
		dfs(G, next, seen)
	}
}

func Code3() {
	g := &Graph{{1, 2, 4}, {0, 3, 4, 8}, {0, 5}, {1, 7, 8}, {0, 1, 8}, {2, 6, 8}, {5, 7}, {3, 6}, {1, 3, 4, 5}}
	dist := bfs(g, 0)
	for _, el := range dist {
		fmt.Println(el)
	}
}

func bfs(G *Graph, s int) []int {
	N := len(*G)
	dist := make([]int, N)
	for i := range dist {
		dist[i] = -1
	}
	que := helper.NewQueue[int]()

	dist[0] = 0
	que.Enqueue(0)

	for !que.Empty() {
		v := que.Dequeue()

		for _, x := range (*G)[v] {
			if dist[x] != -1 {
				continue
			}
			dist[x] = dist[v] + 1
			que.Enqueue(x)

		}
	}
	return dist
}

type C4 struct{}

func (c4 *C4) Dfs(G *Graph, v int, seen []bool) {
	seen[v] = true

	for _, next := range (*G)[v] {
		if seen[next] {
			continue
		}
		c4.Dfs(G, next, seen)
	}
}

func Code4() {
	g := &Graph{{5}, {6, 3}, {5, 7}, {0, 7}, {1, 6}, {}, {7}, {0}}
	seen := make([]bool, len(*g))
	c4 := &C4{}
	c4.Dfs(g, 0, seen)
	for i := range seen {
		if seen[i] {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
