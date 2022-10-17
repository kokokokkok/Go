package main

import (
	"fmt"
	"math"
	"sort"
)

func A_OnAndOff(s, t, x int) {
	if t == 0 {
		t = 24
	}
	if s >= t && t >= x {
		fmt.Println("Yes")
	} else if x >= s && t >= x {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func B_tool_transchack(A_tf *[]bool, A *[]int, next *int, all *int) bool {
	if !(*A_tf)[*next] {
		(*A_tf)[*next] = true
		*next = (*A)[*next] - 1
		*all += 1
		//fmt.Println(*next)
		return true
	} else {
		return false
	}
}

func B_TakahashisSelect(n, x int, A []int) {
	var A_tf = make([]bool, n)
	var all = 0
	//ret := 0
	//fmt.Println(A_tf[2])
	x = x - 1
	for {
		if !B_tool_transchack(&A_tf, &A, &x, &all) {
			break
		}
	}
	fmt.Println("all :", all)
}

func C_FinalDay(n, k int, P [][]int) {
	var st []int
	od := make([]int, n)
	for i := 0; i < n; i++ {
		score := 0
		//fmt.Println(score)
		for j := 0; j < 3; j++ {
			score += P[i][j]
		}
		st = append(st, score)
	}
	copy(od, st)
	sort.Sort(sort.Reverse(sort.IntSlice(st)))
	fmt.Println(st, od)

	for i := 0; i < n; i++ {
		if od[i]+300 >= st[k-1] {
			fmt.Println(("Yes"))
		} else {
			fmt.Println(("No"))
		}
	}

}

func D_LinearProbing(q int, TX [][]int) {
	n := int(math.Pow(2, 20))
	A := make([]int, n, n)

	for i := 0; i < q; i++ {
		if TX[i][0] == 1 {
			h := TX[i][1]
			for {
				if A[h%n] == 0 {
					break
				} else {
					h++
				}
			}
			A[h%n] = TX[i][1]
		} else if TX[i][0] == 2 {
			println(A[TX[i][1]%n])
		}

	}
}
func main() {
	fmt.Print("\nA_Problem\n")
	A_OnAndOff(7, 20, 12)
	A_OnAndOff(20, 7, 12)
	A_OnAndOff(23, 0, 23)
	fmt.Print("\nB_Problem\n")
	var b1 = []int{3, 1, 1, 2}
	B_TakahashisSelect(4, 2, b1)
	var b2 = []int{7, 11, 10, 1, 7, 20, 14, 2, 17, 3, 2, 5, 19, 20, 8, 14, 18, 2, 10, 10}
	B_TakahashisSelect(20, 12, b2)
	fmt.Print("\nC_Problem\n")
	var c1 = [][]int{{178, 205, 132}, {112, 220, 96}, {36, 64, 20}}
	C_FinalDay(3, 1, c1)
	var c2 = [][]int{{300, 300, 300}, {200, 200, 200}}
	C_FinalDay(2, 1, c2)
	var c3 = [][]int{{127, 235, 78}, {192, 134, 298}, {28, 56, 42}, {96, 120, 250}}
	C_FinalDay(4, 2, c3)
	fmt.Print("\nD_Problem\n")
	var d1 = [][]int{{1, 1048577}, {1, 1}, {2, 2097153}, {2, 3}}
	D_LinearProbing(4, d1)
}
