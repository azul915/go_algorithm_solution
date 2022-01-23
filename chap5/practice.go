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

func Code5_3() {
	h := []float64{2, 9, 4, 5, 1, 6, 10}
	N := len(h)

	dp := []float64{0}
	for idx := 1; idx < N; idx++ {
		dp = append(dp, float64(math.Inf(0))-1)
	}

	for idx := 1; idx < N; idx++ {
		chmin(&dp[idx], dp[idx-1]+math.Abs(h[idx]-h[idx-1]))
		if idx > 1 {
			chmin(&dp[idx], dp[idx-2]+math.Abs(h[idx]-h[idx-2]))
		}
	}

	fmt.Printf("dp[N-1]: %v\n", dp[N-1])
	fmt.Printf("db: %v\n", dp)
}