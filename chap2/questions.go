package chap2

import (
	"fmt"
	"math"
)

// p.64
// 2.4
func Question2_4(k int, max int) {

	targetAge := max - 2
	left := 0
	// 2^k未満なので1引く
	right := max - 1
	count := 0

	for left <= right {
		mid := left + (right-left)/2
		count++
		if targetAge < mid {
			right = mid - 1
			fmt.Printf("left: %v, mid: %v right: %v\n", left, mid, right)
		} else if mid == targetAge {
			fmt.Printf("count=%v, k=%v\n", count, k)
			break
		} else /* mid < targetAge */ {
			fmt.Printf("left: %v, mid: %v right: %v\n", left, mid, right)
			left = mid + 1
		}
	}
}

func TestQuestion2_4() {

	for chance := 1; chance <= 4; chance++ {
		m := int(math.Pow(float64(2), float64(chance)))
		Question2_4(chance, m)
	}
}
