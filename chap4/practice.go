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

func fibo(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibo(n-1) + fibo(n-2)
}

func Code4_5() {
	fmt.Printf("fibo(%v) = %v\n", 6, fibo(6))
}

func fib(n int) int {

	fmt.Printf("fib(%v) を呼び出しました\n", n)

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	result := fib(n-1) + fib(n-2)
	fmt.Printf("%v項目 = %v\n", n, result)
	return result
}

func Code4_6() {
	fib(6)
}

// forでのDPによる50までのフィボナッチ数列
func Code4_7() {
	F := [50]int{0, 1}

	for N := 2; N < 50; N++ {
		F[N] = F[N-1] + F[N-2]
		fmt.Printf("%d 項目: %d\n", N, F[N])
	}
}

var memo [50]int

func fibdp(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	if memo[n] != -1 {
		return memo[n]
	}

	memo[n] = fibdp(n-1) + fibdp(n-2)
	return memo[n]
}

// メモ化再帰でのフィボナッチ数列
func Code4_8() {

	// メモ化用配列を-1で初期化する
	for i := 0; i < 50; i++ {
		memo[i] = -1
	}

	fibdp(49)

	// memo[0], ... , memo[49]に答えが格納されている
	for N := 2; N < 50; N++ {
		fmt.Printf("%d 項目: %d\n", N, memo[N])
	}
}

func partlySum(i int, w int, a *[]int) bool {

	// base case
	if i == 0 {
		if w == 0 {
			return true
		} else {
			return false
		}
	}

	// a[i-1]を選ぶ場合
	if partlySum(i-1, w, a) {
		return true
	}

	// a[i-1]を選ばない場合
	if partlySum(i-1, w-(*a)[i-1], a) {
		return true
	}

	return false
}

func Code4_9() {
	N, W := 4, 14
	a := []int{3, 2, 6, 5}

	if partlySum(N, W, &a) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
