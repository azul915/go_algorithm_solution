package chap5

import (
	"fmt"
	"math"
)

func Code5_1() {
	h := []int{2, 9, 4, 5, 1, 6, 10}
	N := len(h)

	var dp []int
	for idx := 0; idx < N; idx++ {
		dp = append(dp, int(math.Inf(0))-1)
	}
	dp[0] = 0

	for idx := 1; idx < N; idx++ {
		if idx == 1 {
			dp[idx] = int(math.Abs(float64(h[idx] - h[idx-1])))
		} else {
			dp[idx] = int(math.Min(float64(int(float64(h[idx-1]))+int(math.Abs(float64(h[idx]-h[idx-1])))), math.Abs(float64(int(float64(h[idx-2]))+int(float64(h[idx]-h[idx-2]))))))
		}
	}

	fmt.Printf("dp[N-1]: %v\n", dp[N-1])
	fmt.Printf("dp: %v\n", dp)

}
