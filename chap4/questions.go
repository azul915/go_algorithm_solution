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
