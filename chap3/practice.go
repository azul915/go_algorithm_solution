package chap3

import (
	"fmt"
)

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
