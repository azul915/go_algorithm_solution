package chap4

import (
	"fmt"
)

// p.542
// 4.1
// トリボナッチ数列の第N項の値を求める再帰関数の設計
func trib(n int) int {

	if n == 0 || n == 1 {
		return 0
	}

	if n == 2 {
		return 1
	}

	return trib(n-1) + trib(n-2) + trib(n-3)
}

func Question4_1() {
	// 第0項~第10項を求める
	k := 10
	for N := 0; N < k; N++ {
		fmt.Printf("%d, ", trib(N))
	}
	fmt.Println("...")
}

// p.542
// 4.2
// 4.1をメモ化によって効率化する
// 計算量 O(N)
var m [10]int

func tribWithMemo(n int) int {

	if n == 0 || n == 1 {
		return 0
	}

	if n == 2 {
		return 1
	}

	if m[n] != -1 {
		return m[n]
	}

	m[n] = tribWithMemo(n-1) + tribWithMemo(n-2) + tribWithMemo(n-3)
	return m[n]
}

func Question4_2() {

	k := 10
	m[0], m[1], m[2] = 0, 0, 1
	// メモ用配列init
	for N := 3; N < k; N++ {
		m[N] = -1
	}

	for N := 0; N < k; N++ {
		fmt.Printf("%d, ", tribWithMemo(N))
	}
	fmt.Println("...")
}

// p.542
// 十進法表記で各桁の値が7,5,3のいずれかであり、
// かつ7,5,3がいずれも一度以上は登場する整数を「753数」と呼ぶとき、
// 正の整数K以下の「753数」が何個あるかを求める
// Kの桁数をdとしてO(3^d)程度の計算量は許容できるものとする
func func753(sft int, limit int, cnt *int) {

	if limit < sft {
		return
	}
	if 0 < sft {
		(*cnt)++
	}

	// 10倍して1の位に7,5,3のいずれかを付与する
	func753(sft*10+7, limit, cnt)
	func753(sft*10+5, limit, cnt)
	func753(sft*10+3, limit, cnt)
}

func Question4_5() {
	K := 10

	count := 0
	func753(0, K, &count)

	fmt.Printf("count: %d\n", count)
}

// p.542
// 部分和問題に対する再帰関数を用いる計算量O(2^N)のコード4.9に対し、
// メモ化して、O(NW)の計算量で動作するようにする
var mem [5][15]int

func partlySumWithMemo(i int, w int, a *[]int) int {

	// base case
	if i == 0 {
		if w == 0 {
			return 1
		} else {
			return 0
		}
	}

	if mem[i][w] != -1 {
		return mem[i][w]
	}

	if partlySumWithMemo(i-1, w, a) == 1 {
		mem[i][w] = 1
		return mem[i][w]
	}

	if partlySumWithMemo(i-1, w-(*a)[i-1], a) == 1 {
		mem[i][w] = 1
		return mem[i][w]
	}

	mem[i][w] = 0
	return mem[i][w]
}

func Question4_6() {
	N, W := 4, 14
	a := []int{3, 2, 6, 5}

	// memoで管理、初期化しておく
	for i := 0; i < N+1; i++ {
		for j := 0; j < W+1; j++ {
			mem[i][j] = -1
		}
	}

	if partlySumWithMemo(N, W, &a) == 1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
