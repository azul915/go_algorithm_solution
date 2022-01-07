package chap3

import (
	"fmt"
	"math/rand"
)

// p.83
// 3.1
func Question3_1() {
	a := []int{1, 2, 3, 2, 4, 5, 6, 2, 7, 8, 9, 10}
	N := len(a)
	v := 2
	foundId := -1

	for i := 0; i < N; i++ {
		if a[i] == v {
			foundId = i
		}
	}
	// 最大のものが格納されることを確認
	fmt.Println("foundId shoud be 7.")
	fmt.Printf("foundId: %v", foundId)
}

// p.83
// 3.2
// N個の整数a0, a1, ..., aN-1のうち、整数値vが何個含まれるかを求めるO(N)のアルゴリズム
func Question3_2() {
	a := []int{1, 2, 3, 2, 4, 5, 6, 2, 7, 8, 9, 10}
	N := len(a)
	v := 2
	cnt := 0

	for i := 0; i < N; i++ {
		if a[i] == v {
			cnt++
		}
	}
	fmt.Printf("cnt: %v", cnt)
}

// p.83
// 3.3
// N (>=2)個の相異なる整数のうち、2番目に小さい値を求めるO(N)のアルゴリズム
func Question3_3() {
	a := []int{}
	for i := 0; i < 10; i++ {
		a = append(a, rand.Intn(100))
	}
	fmt.Printf("a: %v\n", a)

	N := len(a)
	min := 123456789
	second := min - 1

	for i := 0; i < N; i++ {
		if a[i] < min {
			second = min
			min = a[i]
		} else if a[i] < second {
			second = a[i]
		}
		fmt.Printf("min: %v, second: %v\n", min, second)
	}
}
