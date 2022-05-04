package chap10

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Code10.1

type (
	Graph [][]int
)

var (
	G Graph
)

func Code3() {
	f, _ := os.Open("./chap10/input_3.tsv")
	defer f.Close()
	r := csv.NewReader(f)
	r.Comma = '\t'

	fl, _ := r.Read()
	// 頂点数
	N, _ := strconv.Atoi(fl[0])
	// 辺数
	M, _ := strconv.Atoi(fl[1])

	// flag(undirected graph)
	undirected := true
	G = make(Graph, N)
	for i := 0; i < M-1; i++ {
		r, err := r.Read()
		if err == io.EOF {
			break
		}
		a, _ := strconv.Atoi(r[0])
		b, _ := strconv.Atoi(r[1])
		G[a] = append(G[a], b)

		// undirected graph
		if undirected {
			G[b] = append(G[b], a)
		}
	}
	fmt.Printf("G: %v\n", G)
}

type (
	EdgeGraph [][]Edge
	Edge      struct {
		To   int // 頂点隣接番号
		W    int // 重み
		Next *Edge
	}
)

var (
	EG EdgeGraph
)

func Code4() {
	f, _ := os.Open("./chap10/input_4.tsv")
	defer f.Close()
	r := csv.NewReader(f)
	r.Comma = '\t'

	fl, _ := r.Read()
	N, _ := strconv.Atoi(fl[0])
	M, _ := strconv.Atoi(fl[1])

	EG = make(EdgeGraph, N)
	for i := 0; i < M-1; i++ {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		a, _ := strconv.Atoi(record[0])
		b, _ := strconv.Atoi(record[1])
		w, _ := strconv.Atoi(record[2])
		EG[a] = append(EG[a], Edge{To: b, W: w})
	}
	fmt.Println(EG)
}

func Code5() {
	type (
		Heap []int
	)

	var heap Heap
	push := func(x int) {
		// まずはheapの最後尾に入れる
		heap = append(heap, x)

		// 最後尾の添字を求めて、その葉の親ノードの添字を求める
		i := len(heap) - 1

		// 親子関係が崩れているので、ヒープの状態になるまで繰り返す
		for 0 < i {
			p := (i - 1) / 2
			// 挿入する値が親の値より小さいときヒープの状態を満たしていると言える
			if x <= heap[p] {
				break
			}
			// 親子入れ替えを行うために、親の値を子に代入
			heap[i] = heap[p]
			// 自分は親の位置に上がる
			i = p
		}
		// ヒープの状態を満たしているときに、元々挿入する値を入れる
		heap[i] = x
	}

	top := func() int {
		if len(heap) != 0 {
			return heap[0]
		}
		return -1
	}

	// 根の値を消してとりあえず、最後尾を根にしてから、ヒープの状態に戻す方針
	pop := func() {
		// 空のとき
		if len(heap) == 0 {
			return
		}

		x := heap[len(heap)-1]
		i := 0
		for 2*i+1 < len(heap) {
			// 子頂点それぞれの添字を算出して、子頂点の大きな方をchild1として採用
			child1, child2 := 2*i+1, 2*i+2
			if child2 < len(heap) && heap[child1] < heap[child2] {
				child1 = child2
			}
			// 子 <= 親のとき終了
			if heap[child1] <= x {
				break
			}
			// 子頂点の値を親に入れる
			heap[i] = heap[child1]
			// 子頂点に降りる
			i = child1
		}
		// ヒープを満たしたときの添字
		heap[i] = x
	}

	push(5)
	push(3)
	push(7)
	push(1)
	fmt.Println(top())
	pop()
	fmt.Println(top())
	push(11)
	fmt.Println(top())
}
