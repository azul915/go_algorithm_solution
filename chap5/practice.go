package chap5

import (
	"fmt"
	"math"
)

func Code5_1() {
	h := []float64{2, 9, 4, 5, 1, 6, 10}
	N := len(h)

	dp := []float64{0}
	for idx := 1; idx < N; idx++ {
		dp = append(dp, float64(math.Inf(0))-1)
	}

	for idx := 1; idx < N; idx++ {
		if idx == 1 {
			dp[idx] = math.Abs(h[idx] - h[idx-1])
		} else {
			dp[idx] = math.Min(dp[idx-1]+math.Abs(h[idx]-h[idx-1]), dp[idx-2]+math.Abs(h[idx]-h[idx-2]))
		}
	}

	fmt.Printf("dp[N-1]: %v\n", dp[N-1])
	fmt.Printf("dp: %v\n", dp)

}

// Code5.2
func chmin(a *float64, b float64) {
	if *a > b {
		*a = b
	}
}

// Frog問題-貰う遷移方式
func Code5_3() {
	h := []float64{2, 9, 4, 5, 1, 6, 10}
	N := len(h)

	// init
	dp := []float64{0}
	for idx := 1; idx < N; idx++ {
		dp = append(dp, math.Inf(0)-1)
	}

	for idx := 1; idx < N; idx++ {
		fmt.Printf("idx: %v\n", idx)
		fmt.Printf("dp: %v\n", dp)

		chmin(&dp[idx], dp[idx-1]+math.Abs(h[idx]-h[idx-1]))

		fmt.Printf("dp: %v\n", dp)

		if idx > 1 {
			chmin(&dp[idx], dp[idx-2]+math.Abs(h[idx]-h[idx-2]))
		}
		fmt.Printf("dp: %v\n", dp)
	}

	fmt.Printf("dp[N-1]: %v\n", dp[N-1])
	fmt.Printf("dp: %v\n", dp)
}

// Frog問題-配る遷移方式
func Code5_4() {
	h := []float64{2, 9, 4, 5, 1, 6, 10}
	N := len(h)

	// init
	dp := []float64{0}
	for idx := 1; idx < N; idx++ {
		dp = append(dp, math.Inf(0)-1)
	}

	for idx := 0; idx < N; idx++ {
		fmt.Printf("idx: %v\n", idx)
		fmt.Printf("dp: %v\n", dp)
		if idx+1 < N {
			chmin(&dp[idx+1], dp[idx]+math.Abs(h[idx+1]-h[idx]))
		}
		fmt.Printf("dp: %v\n", dp)
		if idx+2 < N {
			chmin(&dp[idx+2], dp[idx]+math.Abs(h[idx+2]-h[idx]))
		}
		fmt.Printf("dp: %v\n", dp)
	}

	fmt.Printf("dp[N-1]: %v\n", dp[N-1])
	fmt.Printf("dp: %v\n", dp)
}

// Frog問題-再帰による単純な全探索
// rec(i): 足場0から足場iまで至るまでの最小コスト
// dp: [0 7 2 3 5 4 8]

// func chminInt(a *int, b int) {
// 	if *a < b {
// 		*a = b
// 	}
// }

var h = []float64{2, 9, 4, 5, 1, 6, 10}

func rec(i int) float64 {

	// base case
	if i == 0 {
		return 0
	}

	res := math.Inf(0)

	chmin(&res, rec(i-1)+math.Abs(h[i]-h[i-1]))
	if 1 < i {
		chmin(&res, rec(i-2)+math.Abs(h[i]-h[i-2]))
	}

	return res

}

func Code5_5() {
	fmt.Printf("rec(0): expect=0, actual=%v\n", rec(0))
	fmt.Printf("rec(1): expect=7, actual=%v\n", rec(1))
	fmt.Printf("rec(2): expect=2, actual=%v\n", rec(2))
	fmt.Printf("rec(3): expect=3, actual=%v\n", rec(3))
	fmt.Printf("rec(4): expect=5, actual=%v\n", rec(4))
	fmt.Printf("rec(5): expect=4, actual=%v\n", rec(5))
	fmt.Printf("rec(6): expect=8, actual=%v\n", rec(6))
}
