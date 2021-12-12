package chap1

import (
	"fmt"
	"math"
)

func ResponseFromChap1() string {
	return "This is response from chap1"
}

func hoge() string {
	return "hoge"
}

// 1.1
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

// 1.2 2^6 = 64 < 100 < 2^7 = 128より、100通りを絞るには7回Yes/Noの質問が必要
func CalcCountOfBinarySearch() {
	requiredCounth := 0
	max := float64(100)

	for {
		limit := math.Pow(2, float64(requiredCounth))

		if max < limit {
			fmt.Printf("requiredCounth: %v limit: 2^%v = %v\n", requiredCounth, requiredCounth, limit)
			break
		}
		requiredCounth++
	}

}
