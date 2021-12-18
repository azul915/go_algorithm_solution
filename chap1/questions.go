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

// 1.3

//    2 a
//  x b c
// -------
//  d 3 e
//  f g
// -------
//  h 4 i
func AlphameticsWithBlanksA() {

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

// 1.4

//                a  b  c  d  e  f
//            x         g  h  i  j
//           ---------------------
//                6  6  k  l  m  n
//             6  o  p  q  r  s
//       t  u  6  6  6  v  w
//       x  y  6  z aa  6
//   -----------------------------
//   ab ac ad ae  6  6 af ag ah ai

type I interface {
	billion() int32
	hundMillion() int32
	tenMillion() int32
	million() int32
	hundThousand() int32
	tenThousand() int32
	thousand() int32
	hundsand() int32
	ten() int32
	one() int32
	number() int32
}

type ue struct {
	HundThousand, TenThousand, Thousand, Hundred, Ten, One int32
}

func (u *ue) hundThousand() int32 {
	return int32(100000) * u.HundThousand
}

func (u *ue) tenThousand() int32 {
	return int32(10000) * u.TenThousand
}

func (u *ue) thousand() int32 {
	return int32(1000) * u.Thousand
}

func (u *ue) hundred() int32 {
	return int32(100) * u.Hundred
}

func (u *ue) ten() int32 {
	return int32(10) * u.Ten
}

func (u *ue) one() int32 {
	return u.One
}

func (u *ue) number() int32 {
	return u.hundThousand() + u.tenThousand() +
		u.thousand() + u.hundred() + u.ten() + u.one()
}

type shita struct {
	Thousand, Hundred, Ten, One int32
}

func (s *shita) thousand() int32 {
	return int32(1000) * s.Thousand
}

func (s *shita) hundred() int32 {
	return int32(100) * s.Hundred
}

func (s *shita) ten() int32 {
	return int32(10) * s.Ten
}

func (s *shita) one() int32 {
	return s.One
}

func (s *shita) number() int32 {
	return int32(1000)*s.Thousand + int32(100)*s.Hundred + int32(10)*s.Ten + s.One
}

type answer struct {
	Billion, HundMillion, TenMillion, Million, HundThousand, TenThousand, Thousand, Hundred, Ten, One int32
}

func (a *answer) number() int32 {
	return int32(1000000000)*a.Billion + int32(100000000)*a.HundMillion +
		int32(10000000)*a.TenMillion + int32(1000000)*a.Million +
		int32(100000)*a.HundThousand + int32(10000)*a.TenThousand +
		int32(1000)*a.Thousand + int32(100)*a.Hundred + int32(10)*a.Ten + a.One
}

func (a *answer) billion() int32 {
	return int32(1000000000) * a.Billion
}

func (a *answer) hundMillion() int32 {
	return int32(100000000) * a.HundMillion
}

func (a *answer) tenMillion() int32 {
	return int32(10000000) * a.TenMillion
}

func (a *answer) million() int32 {
	return int32(1000000) * a.Million
}

func (a *answer) hundThousand() int32 {
	return int32(100000) * a.HundThousand
}

func (a *answer) tenThousand() int32 {
	return int32(10000) * a.TenThousand
}

func (a *answer) thousand() int32 {
	return int32(1000) * a.Thousand
}

func (a *answer) hundred() int32 {
	return int32(100) * a.Hundred
}

func (a *answer) ten() int32 {
	return int32(10) * a.Ten
}

func (a *answer) one() int32 {
	return a.One
}

//                a  b  c  d  e  f
//            x         g  h  i  j
//           ---------------------
//                6  6  k  l  m  n
//             6  o  p  q  r  s
//       t  u  6  6  6  v  w
//       x  y  6  z aa  6
//   -----------------------------
//   ab ac ad ae  6  6 af ag ah ai

func dfs(upper ue, down shita) {

	if upper.TenThousand*down.One != int32(6) ||
		upper.HundThousand*down.One != int32(6) ||
		upper.TenThousand*down.Ten != int32(6) ||
		upper.Hundred*down.Hundred != int32(6) ||
		upper.Thousand*down.Hundred != int32(6) ||
		upper.TenThousand*down.Hundred != int32(6) ||
		upper.One*down.Thousand != int32(6) ||
		upper.Thousand*down.Thousand != int32(6) {

	}

}

func AlphameticsWithBlanksB() {

	u := ue{HundThousand: 1, TenThousand: 0, Thousand: 0, Hundred: 0, Ten: 0, One: 0}
	s := shita{Thousand: 1, Hundred: 0, Ten: 0, One: 0}
	_ = answer{Billion: 1, HundMillion: 0, TenMillion: 0, Million: 0,
		HundThousand: 0, TenThousand: 0, Thousand: 0, Hundred: 0, Ten: 0, One: 0}

	dfs(u, s)

}

// 1.5

// 1.6
// - ダイクストラ法(公共交通機関乗換アプリ)
