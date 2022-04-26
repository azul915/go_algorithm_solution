package chap6

import (
	"bufio"
	"fmt"
	"os"
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

func binarySearchB() int {

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

func Code6_5() {
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
