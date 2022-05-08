package chap12

import "fmt"

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
