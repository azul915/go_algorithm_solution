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
