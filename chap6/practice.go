package chap6

import (
	"algorithm_solution/pkg/helper"
	"bufio"
	"fmt"
	"math"

	"os"
	"sort"
)

func Code6_1() {

	fmt.Println(binarySearch(10))
	fmt.Println(binarySearch(3))
	fmt.Println(binarySearch(39))
	fmt.Println(binarySearch(-100))
	fmt.Println(binarySearch(9))
	fmt.Println(binarySearch(100))
}

func binarySearch(key int) int {

	a := []int{3, 5, 8, 10, 14, 17, 21, 39}

	left, right := 0, len(a)-1
	for left <= right {
		mid := (left + right) / 2
		if a[mid] == key {
			return mid
		} else if a[mid] < key {
			left = mid + 1
		} else /* if key < a[mid] */ {
			right = mid - 1
		}
	}
	return -1
}

func P(x int) bool {
	return true
}

func Code6_2() int {

	var left, right int
	for 1 < right-left {
		mid := left + (right-left)/2
		if P(mid) {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}

func Code6_3() {
	fmt.Println("Start Game")
	left, right := 20, 36

	for 1 < right-left {
		mid := left + (right-left)/2
		fmt.Printf("Is the age less than %v ? (yes/no)\n", mid)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		ans := scanner.Text()

		if ans == "yes" {
			right = mid
		} else /* if ans == "no" */ {
			left = mid
		}
	}

	fmt.Printf("The age is %d\n", left)
}

func Code6_4() {
	min := math.MaxInt
	K := 3
	a := []int{3, 5, 8, 10, 14, 17, 21, 39}

	b := []int{4, 2, 9, 14, 7, 18, 20, 29}
	sort.Ints(b)

	for i := 0; i < len(a); i++ {
		iter := helper.LowerBound(K-a[i], b)
		val := b[iter]
		if a[i]+val < min {
			min = a[i] + val
		}
	}

	fmt.Println(min)
}

func Code6_5() {
	N := 4
	h := []int{5, 12, 14, 21}
	s := []int{6, 4, 7, 2}
	left, right := 0, math.MaxInt

	for 1 < right-left {
		mid := (left + right) / 2
		ok := true
		t := make([]int, N)
		for i := 0; i < N; i++ {
			if mid < h[i] {
				ok = false
			} else {
				t[i] = (mid - h[i]) / s[i]
			}
		}
		sort.Ints(t)
		for i := 0; i < N; i++ {
			if t[i] < i {
				ok = false
			}
		}
		if ok {
			right = mid
		} else {
			left = mid
		}
	}
	fmt.Println(right)
}
