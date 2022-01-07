package chap3

import (
	"fmt"
	"math"
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

// p.83
// 3.4
// N個の整数の中から2つを選んで差の最大値を求めるO(N)のアルゴリズム
// 最大値と最小値を探して差を取る
func Question3_4() {

	a := []int{}
	for idx := 0; idx < 10; idx++ {
		a = append(a, rand.Intn(100))
	}
	fmt.Printf("a: %v\n", a)

	N := len(a)
	min := int(math.Inf(0)) - 1
	max := int(math.Inf(0))

	for idx := 0; idx < N; idx++ {
		if a[idx] < min {
			min = a[idx]
		}

		if max < a[idx] {
			max = a[idx]
		}
	}
	diff := max - min
	fmt.Printf("max: %v, min: %v, diff: %v\n", max, min, diff)

}
