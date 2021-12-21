package chap2

import (
	"fmt"
	"math"
)

func Foo() {
	fmt.Println("")
}

func calcDist(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
}

// p52, code2.4
func Code2_4() float64 {
	var N int = 3
	var x []float64 = []float64{1, 3, 5}
	var y []float64 = []float64{2, 1, 6}

	// 求める値について十分大きな値での初期化
	minimumDist := math.Inf(0)

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			distIandJ := calcDist(x[i], y[i], x[j], y[j])
			if distIandJ < minimumDist {
				minimumDist = distIandJ
			}
		}
	}
	return minimumDist
}
