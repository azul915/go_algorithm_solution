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
	for N := 0; N < 10; N++ {
		fmt.Printf("%d, ", trib(N))
	}
	fmt.Println("...")
}
