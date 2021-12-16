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

//    2 a
//  x b c
// -------
//  d 3 e
//  f g
// -------
//  h 4 i
func AlphameticsWithBlanks() {

	for a := 0; a <= 9; a++ {
		for b := 0; b <= 9; b++ {
			for c := 0; c <= 9; c++ {

				ue := (20 + a) * c
				if ue < 129 || ue > 939 {
					continue
				}

				// 10の位の桁が3
				if ue%100/10 != 3 {
					continue
				}

				shita := (20 + a) * (10 * b)
				if shita < 100 {
					continue
				}

				sum := ue + shita
				if sum > 950 || sum < 140 {
					continue
				}

				// 10の位の桁が4
				if sum%100/10 != 4 {
					continue
				}

				d := ue / 100
				e := ue - (d*100 + 30)

				f := shita / 100
				g := shita % 100 / 10

				h := sum / 100
				i := sum - (h*100 + 40)

				fmt.Printf("   2 %v\n", a)
				fmt.Printf(" x %v %v\n", b, c)
				fmt.Println("------")
				fmt.Printf(" %v 3 %v\n", d, e)
				fmt.Printf(" %v %v\n", f, g)
				fmt.Println("------")
				fmt.Printf(" %v 4 %v", h, i)
			}
		}
	}
}
