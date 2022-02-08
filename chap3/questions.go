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

// p.83
// 3.5 AtCoder Beginner Contest 081 B - Shift Only
// https://atcoder.jp/contests/abc081/tasks/abc081_b
// N個の正の整数に対してN個の整数がすべて偶数ならば2で割った値に置き換える操作を
// 操作が行えなくなるまで繰り返したとき、何回できるかを求める。
// N個の正の整数に対して、奇数になるまで2で割る操作を行い、その操作回数の最小値を取る
func Question3_5() {

	a := []int{382253568, 723152896, 37802240, 379425024, 404894720, 1471526144}
	min := int(math.Inf(0)) - 1
	N := len(a)

	for idx := 0; idx < N; idx++ {
		count := 0
		for a[idx]%2 == 0 {
			a[idx] /= 2
			count++
		}

		if count < min {
			min = count
		}

	}

	fmt.Printf("min: %v", min)
}

// p.83
// 3.6 AtCoder Beginner Contest 051 B - Sum of Three Integers
// https://atcoder.jp/contests/abc051/tasks/abc051_b
// 2つの正の整数K, Nが与えられる。0 <= X,Y,Z <= Kを満たす整数(X,Y,Z)の組であって
// X + Y + Z = Nを満たすものが何通りあるかを求めるO(N^2)のアルゴリズム
func Question3_6() {

	count := 0
	K, N := 2, 2
	for X := 0; X <= K; X++ {
		for Y := 0; Y <= K; Y++ {
			Z := N - X - Y
			if 0 <= Z && Z <= K {
				fmt.Printf("X: %v, Y: %v, Z: %v\n", X, Y, Z)
				count++
			}
		}
	}
	fmt.Printf("count: %v", count)
}

// p.83
// 3.7 AtCoder Beginner Contest 045 C - たくさんの数式
// https://atcoder.jp/contests/abc045/tasks/arc061_a
// 各桁の値が1以上9以下の数値のみである整数とみなせるような、長さNの文字列が与えられる。
// この文字列の中で、文字と文字の間のうちのいくつかの場所に「+」を入れることができる。
// 1つも入れなくても構わないが、「+」が連続してはいけない。このようにしてできる、全ての文字列を
// 数値とみなして、総和を計算するO(2N)のアルゴリズム
// func Question3_7() {

// 	S := "251"
// 	N := len(a)

// 	for i := 0; i < (1 << (N - 1)); i++ {
// 		tmp := 0
// 		for j := 0; j < N-1; j++ {
// 			tmp *= 10
// 			tmp += S[i:i+1] - '0'
// 			if i&(1<<j) == 0 {
// 				fmt.Printf("i: %v, j: %v\n", i, j)
// 			}
// 		}
// 	}
// }
