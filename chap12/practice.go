package chap12

import (
	"fmt"
)

func Code1() {
	sl := []int{4, 1, 5, 3, 2}
	insertionSort(sl)
	fmt.Println(sl)
}

// in-place
// stable
func insertionSort(a []int) {

	for i, v := range a {
		j := i
		// 後ろから検証、添字jから見て手前(j-1)のときの値が大きいとき、swap
		for j = i; 0 < j; j-- {
			if v < a[j-1] {
				a[j] = a[j-1]
			} else {
				break
			}
		}
		// 最後外側のforで検証している値について添字jの位置に入れる
		a[j] = v
	}
}

func Code2() {
	sl := []int{4, 1, 5, 3, 2}
	mergeSort(sl, 0, len(sl))
	fmt.Println(sl)
}

func mergeSort(a []int, left, right int) {
	if right-left < 2 {
		return
	}
	mid := left + (right-left)/2

	mergeSort(a, left, mid)
	mergeSort(a, mid, right)

	buf := []int{}
	for i := left; i < mid; i++ {
		buf = append(buf, a[i])
	}
	for i := right - 1; mid <= i; i-- {
		buf = append(buf, a[i])
	}

	lidx := 0
	ridx := len(buf) - 1
	for i := left; i < right; i++ {
		if buf[lidx] <= buf[ridx] {
			a[i] = buf[lidx]
			lidx++
		} else {
			a[i] = buf[ridx]
			ridx--
		}
	}
}

func Code3() {
	sl := []int{4, 1, 5, 3, 2}
	quickSort(sl, 0, len(sl))
	fmt.Println(sl)
}

func quickSort(a []int, left, right int) {
	if right-left < 2 {
		return
	}
	piv_idx := left + (right-left)/2
	pivot := a[piv_idx]
	a[piv_idx], a[right-1] = a[right-1], a[piv_idx]
	i := left
	for j := left; j < right-1; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[right-1] = a[right-1], a[i]

	quickSort(a, left, i)
	quickSort(a, i+1, right)
}

// func Code4() {
// 	sl := []int{4, 1, 5, 3, 2}
// 	heapSort(sl)
// 	fmt.Println(sl)
// }

// func heapSort(a []int) {
// 	N := len(a)
// 	for i := N/2 - 1; 0 <= N; i-- {
// 		Heapify(a, i, N)
// 	}

// 	for i := N - 1; 0 < i; i-- {
// 		a[0], a[i] = a[i], a[0]
// 		Heapify(a, 0, i)
// 	}
// }

// func Heapify(a []int, i, N int) {
// 	child1 := i*2 + 1
// 	if N <= child1 {
// 		return
// 	}
// 	if child1+1 < N && a[child1] < a[child1+1] {
// 		child1++
// 	}
// 	if a[child1] <= a[i] {
// 		return
// 	}
// 	a[i], a[child1] = a[child1], a[i]
// 	Heapify(a, child1, N)
// }

// func Code5() {
// 	sl := []int{4, 1, 5, 3, 2}
// 	bucketSort(sl)
// 	fmt.Println(sl)
// }

// func bucketSort(a []int) {
// 	N := len(a)
// 	MAX := 100000

// 	num := make([]int, MAX)
// 	for i := range num {
// 		num[i] = 0
// 	}
// 	for i := 0; i < N; i++ {
// 		num[a[i]]++
// 	}

// 	sum := make([]int, MAX)
// 	for i := range sum {
// 		sum[i] = 0
// 	}
// 	for v := 1; v < MAX; v++ {
// 		sum[v] = sum[v-1] + sum[v]
// 	}

// 	a2 := make([]int, N)
// 	for i := N - 1; 0 <= i; i-- {
// 		idx := sum[a[i]]
// 		idx--
// 		a2[idx] = a[i]
// 	}
// 	a = a2
// }
