package chap4

import (
	"fmt"
)

func func4_1(N int) int {
	if N == 0 {
		return 0
	}
	return N + func4_1(N-1)
}

func Code4_1() {
	N := 5
	fmt.Printf("func4_1(%v): %v\n", N, func4_1(N))
}

func func4_2(N int) int {
	fmt.Printf("func4_2(%v)を呼び出しました\n", N)

	if N == 0 {
		return 0
	}
	result := N + func4_2(N-1)
	fmt.Printf("%v までの和 = %v\n", N, result)
	return result

}

func Code4_2() {
	N := 5
	func4_2(N)
}

// p.19 ユークリッドの互除法によって最大公約数を求める
func GCD(m int, n int) int {
	if n == 0 {
		return m
	}

	return GCD(n, m%n)
}

func Code4_4() {
	fmt.Printf("GCD(%v, %v) = %v\n", 51, 15, GCD(51, 15))
	fmt.Printf("GCD(%v, %v) = %v\n", 15, 51, GCD(15, 51))
}
