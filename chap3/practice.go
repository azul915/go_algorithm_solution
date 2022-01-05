package chap3

import (
	"fmt"
	"math"
	"strings"
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

// p80 code3.6
// 部分集合{a0, a2, a3, a6}にi番目の要素aiが含まれるかどうかを判定する
// | i | 1 << i   | bit & (1 << i)                        |
// |---|----------|---------------------------------------|
// | 0 | 00000001 | 01001101 & 00000001 = 00000001(true)  |
// | 1 | 00000010 | 01001101 & 00000010 = 00000000(false) |
// | 2 | 00000100 | 01001101 & 00000100 = 00000100(true)  |
// | 3 | 00001000 | 01001101 & 00001000 = 00001000(true)  |
// | 4 | 00010000 | 01001101 & 00010000 = 00000000(false) |
// | 5 | 00100000 | 01001101 & 00100000 = 00000000(false) |
// | 6 | 01000000 | 01001101 & 01000000 = 01000000(true)  |
// | 7 | 10000000 | 01001101 & 10000000 = 00000000(false) |

// N = 5, W = 10 a = {1, 2, 4, 5, 11}の場合には、a0+a2+a3 = 1+4+5 = 10でYesとなる
func UseBitCalcInSubsetSumProblem() {

	W := 10
	a := []int{1, 2, 4, 5, 11}
	N := len(a)
	exists := false

	for bit := 0; bit < (1 << N); bit++ {
		sum := 0
		for idx := 0; idx < N; idx++ {
			// どこにもビットが立たないときは00000000となり、それ以外がどこかにbitが立つときである
			// fmt.Printf("1<<bit: %v\n", 1<<bit)
			fmt.Printf("bit: %v\n", bit)
			fmt.Printf("bit&(1<<bit): %v\n", bit&(1<<bit))
			if bit&(1<<bit) != 0 {
				sum += a[idx]
			}
		}

		if sum == W {
			exists = true
		}
		// fmt.Printf("%v\n", sum)
	}

	if exists {
		fmt.Printf("Yes")
	} else {
		fmt.Printf("No")
	}
}

func BitCheck() {

	var bit int
	bit = 3
	for i := 0; i < 10; i++ {
		// 2は00000010
		// 3は00000011
		if i&bit != 0 {
			fmt.Print("-> ")
		}
		fmt.Printf("i: %v, bin: %v\n", i, convertToBinary(i))
	}
}

func convertToBinary(decimal int) string {
	binstr := ""
	for 0 < decimal {
		surplus := decimal % 2
		binstr = strings.Join([]string{fmt.Sprint(surplus), binstr}, "")
		decimal /= 2
	}
	return binstr
}
