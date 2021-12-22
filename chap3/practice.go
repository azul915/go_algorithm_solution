package chap3

import (
	"fmt"
	"math"
)

// p70 code3.1
func BinarySearchSample(find int32) bool {

	var target []int32 = []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var exist bool

	// i := rangeのように変数を1つとしたときは、sliceや配列のindexが取れる。i, _ := rangeの省略系
	// for i := range target {
	// 	if target[i] == find {
	// 		exist = true
	// 		break
	// 	}
	// }

	for _, v := range target {
		if v == find {
			exist = true
			break
		}
	}

	if exist {
		fmt.Printf("%v is exists in target\n", find)
	} else {
		fmt.Printf("%v is not in target\n", find)
	}
	return exist
}

// p71 code3.2
func BinarySearchWithFoundIdSample(find int32) (bool, int) {

	target := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var exist bool
	foundId := -1
	for i := range target {
		if target[i] == find {
			exist = true
			foundId = i
			break
		}
	}

	if exist {
		fmt.Printf("%v is exists in target (foundId: %v)\n", find, foundId)
	} else {
		fmt.Printf("%v is not in target\n", find)
	}

	return exist, foundId
}

// p72 code3.3
func SearchMinimumNumber() {

	target := []int32{24, 41, 84, 103, 54, 13, 5, 64}
	// big number enough
	min := math.Inf(0)

	for _, v := range target {
		if float64(v) < min {
			min = float64(v)
		}
	}

	fmt.Printf("%v is the smallest number in target", min)
}

// p76 code3.4
func MinimumNumberInTargetPair() {

	var N int32
	var K int32
	var a []int32
	var b []int32
	minValue := math.Inf(0)

	for i := 0; i < int(N); i++ {
		for j := 0; j < int(N); j++ {
			if a[i]+b[j] < K {
				continue
			}

			if float64(a[i]+b[j]) < minValue {
				minValue = float64(a[i] + b[j])
			}
		}
	}

	fmt.Printf("%v is the smallest number in the pair", minValue)
}
