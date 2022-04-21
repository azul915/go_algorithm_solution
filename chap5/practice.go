package chap5

import (
	"fmt"
	"math"
	_ "reflect"
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
		dp = append(dp, math.Inf(0))
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
		dp = append(dp, math.Inf(0))
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

var hh = []float64{2, 9, 4, 5, 1, 6, 10}
var dp = []float64{0}
var N = len(h)
var INF = math.Inf(0)

func recWithMemo(i int) float64 {

	if dp[i] < INF {
		return dp[i]
	}

	// base case
	if i == 0 {
		return 0
	}

	var res = math.Inf(0)
	chmin(&res, recWithMemo(i-1)+math.Abs(hh[i]-hh[i-1]))
	if 1 < i {
		chmin(&res, recWithMemo(i-2)+math.Abs(hh[i]-hh[i-2]))
	}

	dp[i] = res
	return dp[i]

}

func Code5_6() {

	// init
	for idx := 1; idx < N; idx++ {
		dp = append(dp, INF)
	}

	// fmt.Printf("rec(0): expect=0, actual=%v\n", recWithMemo(0))
	// fmt.Printf("rec(1): expect=7, actual=%v\n", recWithMemo(1))
	// fmt.Printf("rec(2): expect=2, actual=%v\n", recWithMemo(2))
	// fmt.Printf("rec(3): expect=3, actual=%v\n", recWithMemo(3))
	// fmt.Printf("rec(4): expect=5, actual=%v\n", recWithMemo(4))
	fmt.Printf("rec(5): expect=4, actual=%v\n", recWithMemo(5))
	// fmt.Printf("rec(6): expect=8, actual=%v\n", recWithMemo(6))

}

// [問題]
// N個の品物があり、i(=0, 1, .., N-1)番目の品物の重さはweight{i}, 価値はvalue{i}で与えられる
// N個の品物から重さの総和がWを超えないようにいくつか選ぶとき、選んだ品物の価値の総和として考えられる最大値を求める

// [定石]
// N個の対象物{0, 1, ..., N-1}に関する問題に対して、最初のi個の対象物{0, 1, ..., i-1}に関する問題を部分問題として考える

// [Note]
// 0, 1, ..., i-1番目の品物からいくつか選んだあとに、i番目の品物を「選ぶ」「選ばない」という2通りの選択肢が存在する = 「各段階においていくつかの選択肢が存在する」
// dp[i][w] : 最初のi個の品物{0, 1, ..., i-1}までの中から重さがwを超えないように選んだときの価値の総和の最大値
// 動的計画法とは考えられる場合をグループごとにまとめるイメージの手法
// 品物がないときは重さも価値もともに0なので、dp[0][w] = 0(w = 0, 1, ..., W)
// dp[i+1][w] (w = 0, 1, ..., W)について場合分けして考える
// A. i番目の品物を選ぶとき
// 選んだときの価値を(i+1, w)とすると、選ぶ前の価値は(i, w-weight[i])と表すこととができる。そこにvalue[i]が加わるので
// chmax(dp[i+1][w], dp[i][w-weight[i]]+value[i]) w - weight[i] >= 0の場合のみ
// B. i番目の品物を選ばないとき
// 重さや価値に変化はないので
// chmax(dp[i+1][w], dp[i][w])

func chmax(a *int, b int) {
	if *a < b {
		*a = b
	}
}

func fmtdegit(num int) string {
	if 99 < num {
		return fmt.Sprintf("%d ", num)
	} else if 9 < num {
		return fmt.Sprintf(" %d ", num)
	} else {
		return fmt.Sprintf("  %d ", num)
	}
}

func Code5_7() {

	d := []struct {
		weight int
		value  int
	}{
		{2, 3}, {1, 2}, {3, 6}, {2, 1}, {1, 3}, {5, 85},
	}
	N, W := len(d), 15

	// DPテーブル定義
	// N:6, W: 15
	// dp 7x16
	dp := make([][]int, N+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}

	for i := 0; i < N; i++ {
		for w := 0; w <= W; w++ {
			// i番目の品物を選ぶとき
			if w-d[i].weight >= 0 {
				chmax(&dp[i+1][w], dp[i][w-d[i].weight]+d[i].value)
			}
			// i番目の品物を選ばないとき
			chmax(&dp[i+1][w], dp[i][w])
		}
	}

	for x := 0; x < len(dp); x++ {
		for y := 0; y < len(dp[x]); y++ {
			fmt.Print(fmtdegit(dp[x][y]))
		}
		fmt.Println("")
	}
}
