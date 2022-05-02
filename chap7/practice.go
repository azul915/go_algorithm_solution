package chap7

import (
	"fmt"
)

func Code7_1() {
	value := []int{500, 100, 50, 10, 5, 1}
	X := 729
	a := []int{1, 2, 0, 2, 1, 4}

	result := 0
	for i := range value {
		add := X / value[i]
		if a[i] < add {
			add = a[i]
		}
		X -= value[i] * add
		result += add
	}
	fmt.Println(result)
}
