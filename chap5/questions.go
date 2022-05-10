package chap5

import (
	// "fmt"
	"algorithm_solution/pkg/helper"
)

// p.154
// 5.1
// https://github.com/drken1215/book_algorithm_solution/blob/master/solutions/chap05.md#51-edpc-c---vacation
func Question5_1() {

	N := 10
	a := [][]int{
		{3, 2, 1}, {3, 2, 1}, {3, 2, 1}, {3, 2, 1}, {3, 2, 1},
		{3, 2, 1}, {3, 2, 1}, {3, 2, 1}, {3, 2, 1}, {3, 2, 1},
	}

	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, 3)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if j == k {
					continue
				}
				helper.ChMax(&dp[i+1][k], dp[i][j]+a[i][k])
			}
		}
	}

	PrintTable(dp)
}
