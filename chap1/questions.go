package chap1

import (
	"fmt"
)

func ResponseFromChap1() string {
	return "This is response from chap1"
}

func hoge() string {
	return "hoge"
}

// Aさんの年齢が22歳、範囲が20歳以上36歳未満だとする
func SearchAgeOfA() int {

	targetAge := 22
	left := 20
	right := 35

	for left <= right {
		mid := left + ((right - left) / 2)
		if mid > targetAge {
			right = mid - 1
			fmt.Printf("left: %v, mid: %v right: %v\n", left, mid, right)
		} else if mid == targetAge {
			fmt.Printf("left: %v, mid: %v right: %v\n", left, mid, right)
			return mid
		} else { // mid < targetAge
			left = mid + 1
			fmt.Printf("left: %v, mid: %v right: %v\n", left, mid, right)
		}
	}
	return -1
}
