package chap7

import (
	"fmt"
	"sort"
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

func Code7_2() {
	jobs := []struct {
		start int
		end   int
	}{
		{start: 9, end: 16},
		{start: 11, end: 15},
		{start: 19, end: 23},
		{start: 10, end: 12},
		{start: 15, end: 18},
		{start: 13, end: 19},
	}
	// 終了時刻が早い順にソートする
	sort.Slice(jobs, func(i, j int) bool { return jobs[i].end < jobs[j].end })

	res := 0
	currentEndTime := 0

	for i := range jobs {
		if jobs[i].start < currentEndTime {
			continue
		}
		res++
		currentEndTime = jobs[i].end
	}
	fmt.Println(res)
}
