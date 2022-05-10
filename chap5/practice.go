package chap5

import (
	"algorithm_solution/pkg/helper"
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
// w - weight[i] >= 0という条件は、wを超えてはならない最大化問題であるということ
// B. i番目の品物を選ばないとき
// 重さや価値に変化はないので
// chmax(dp[i+1][w], dp[i][w]) ちなみに、w = 0のとき、i番目の品物を選ぶことは絶対にない(w = 0なので、重さ0を超えてはならない)

func chmax(a *int, b int) {
	if *a < b {
		*a = b
	}
}

func FmtDegit(num int) string {
	if 99 < num {
		return fmt.Sprintf("%d ", num)
	} else if 9 < num {
		return fmt.Sprintf(" %d ", num)
	} else {
		return fmt.Sprintf("  %d ", num)
	}
}

func PrintTable(table [][]int) {
	for _, iv := range table {
		for _, wv := range iv {
			fmt.Print(FmtDegit(wv))
		}
		fmt.Println("")
	}
	fmt.Println("")
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
	PrintTable(dp)

	for i := 0; i < N; i++ {
		for w := 0; w <= W; w++ {
			// i番目の品物を選ぶとき
			if w-d[i].weight >= 0 {
				chmax(&dp[i+1][w], dp[i][w-d[i].weight]+d[i].value)
			}
			// i番目の品物を選ばないとき
			chmax(&dp[i+1][w], dp[i][w])
		}
		PrintTable(dp)
	}
	PrintTable(dp)
}

// [問題]
// 2つの文字列S, Tが与えられます。Sに以下の3通りの操作を繰り返し施すことでTに変換したいものとします
// そのような一連の操作のうち、操作回数の最小値を求めてください。
// なお、この最小値をSとTとの編集距離と呼びます
// 変更: S中の文字を1つ選んで任意の文字に変更する
// 削除: S中の文字を1つ選んで削除する
// 挿入: Sの好きな箇所に好きな文字を1文字挿入する

// 2つの文字列S, Tの編集距離における
// 片方の文字列Sへの挿入操作は、もう一方のの文字列Tの削除操作と等価のものとして扱えることがポイント

// dp[i][j]: Sの最初のi文字分と、Tの最初のj文字分との間の編集距離
// dp[0][0] = 0: Sの最初の0文字分とTの最初の0文字分がともに空の文字列を表しており、空の文字列同士は特に変更操作を実施することがないから0となる
// 遷移を考えたときに、Sの最初のi文字分とTの最初のj文字分とで、それぞれ最後の1文字をどのように対応したかで場合分けする
// 変更操作(Sのi文字目とTのj文字目とを対応させる)
//     1. 一致するとき: コストなし、つまりchmin(dp[i][j], dp[i-1][j-1])
//     2. 一致しないとき: コスト+1、つまりchmin(dp[i][j], dp[i-1][j-1]+1)
// 削除操作(Sのi文字目を削除)
//     1. コスト+1、つまりchmin(dp[i][j], dp[i-1][j]+1)
// 挿入操作(Tのj文字目を削除)
//     1. コスト+1、つまりchmin(dp[i][j], dp[i][j-1]+1)

func Code5_8() {
	S, T := "logistic", "algorithm"

	dp := make([][]int, len(S)+1)
	for i := range dp {
		dp[i] = make([]int, len(T)+1)
	}

	for si := 0; si <= len(S)+1; si++ {
		for ti := 0; ti < len(T); ti++ {
			if si == 0 && ti == 0 {
				dp[ti][si] = 0
			} else {
				dp[ti][si] = math.MaxInt
			}
		}
	}

	// printTable(dp)

	for si := 0; si <= len(S); si++ {
		for ti := 0; ti <= len(T); ti++ {
			if 0 < si && 0 < ti {
				if string(S[si-1]) == string(T[ti-1]) {
					helper.ChMin(&dp[si][ti], dp[si-1][ti-1])
				} else {
					helper.ChMin(&dp[si][ti], dp[si-1][ti-1]+1)
				}
			}
			if 0 < si {
				helper.ChMin(&dp[si][ti], dp[si-1][ti]+1)
			}
			if 0 < ti {
				helper.ChMin(&dp[si][ti], dp[si][ti-1]+1)
			}
		}
	}
	PrintTable(dp)
	fmt.Println(dp[len(S)][len(T)])
}

// [問題]
// N個の要素が1列に並んでいて、これをいくつかの区間に分割したいものとする
// 各区間[l,r)にはスコアcl,rがついているとする
// KをN以下の正の整数として、K+1個の整数t0,t1,...,tkを0 = t0 < t1 < ... < tk = Nを満たすように取ったとき
// 区間分割[t0,t1),[t1,t2),...,[t(k-1),tka)のスコアを
// Ct0,t1 + Ct1,t2 + ... + Ct(k-1).tk によって定義
// N要素の区間分割の仕方をすべて考えたときの考えられるスコアの最小値を求める

// N = 10, K = 4, t = (0, 3, 7, 8, 10)
// e.g) C0,3 + C3,7 + C7,8 + C8,10
// | o o o | o o o o | o | o o |
// Nは分割対象の数
// Kは分割区間の数
// tは分割方法

// dp[i]: 区間[0,i)について、いくつかの区間に分割する最小コスト
// 初期条件 dp[0] = 0
// 区間[0,i)を分割する方法のうち、最後に区切る場所がどこであったかで場合分け
// 最後に区切る位置がj(= 0,1,...,i-1)であったとき、区間[0,i)の分割は、
// 区間[0,j)の分割に対して新たに区間[j,i)を追加したもの
// とみなすことができる
func Code5_9() {
	N, _ := 10, 4

	c := make([][]int, N+1)
	fmt.Println(c)
	for idx := range c {
		c[idx] = make([]int, N+1)
	}
	fmt.Println(c)

	// DPテーブル
	dp := make([]int, N+1)
	for idx := range dp {
		if idx == 0 {
			dp[idx] = 0
		} else {
			dp[idx] = math.MaxInt
		}
	}

	// fmt.Println(dp)

	for i := 0; i <= N; i++ {
		for j := 0; j <= i; j++ {
			helper.ChMin(&dp[i], dp[j]+c[j][i])
		}
	}

}
