package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func A_FirstGrid(s1 string, s2 string) {
	//fmt.Println(string(s1[0]), s2)
	if string(s1[0]) == "#" && string(s1[1]) == "." && string(s2[0]) == "." && string(s2[1]) == "#" || string(s1[0]) == "." && string(s1[0]) == "#" && string(s2[0]) == "." && string(s2[1]) == "#" {
		fmt.Print("No\n")
	} else {
		fmt.Print("Yes\n")
	}
}

func B_HardCalculation(a string, b string) {
	for i := 0; i < len(a); i++ {
		na, _ := strconv.Atoi(string(a[i]))
		nb, _ := strconv.Atoi(string(b[i]))
		//fmt.Print(na, nb," ")
		if n := na + nb; n > 10 {
			//fmt.Print(na, " ", nb, " ", n, " ", "Hard\n")
			return
		}
	}
	fmt.Print("Easy\n")
}

func C_Cheese(n int, w int, AB [][]int) {
	var max float64 = 0
	for bit := 0; bit < (1 << n); bit++ {
		score := 0
		g := 0
		for i := 0; i < n; i++ {
			//fmt.Println(g+AB[i][1])
			if (bit&(1<<i) == (1 << i)) && ((g + AB[i][1]) <= w) {
				g += AB[i][1]
				score += AB[i][0] * AB[i][1]
				fmt.Println("g:", g, " score:", score)
			}
		}
		max = math.Max(float64(max), float64(score))
	}
	fmt.Println((max))
}

func C_Cheese_New(n int, w int, AB [][]int) {
	var score int = 0
	var g int = 0
	sort.Slice(AB, func(i, j int) bool { return AB[i][0] > AB[j][0] })
	//fmt.Print(AB,"\n")
	s := 0
	for i := 0; i < 100000; i++ {
		if (g+1 <= w) && AB[s][1] > 0 {
			score += AB[s][0]
			g += 1
			AB[s][1] -= 1
			//fmt.Println("g:", g, " score:", score)
			//fmt.Println("[1]:", AB[s][1]," s:",s)
			//fmt.Println("[1]:", AB[s+1][1],"ax",s+1)
		} else if AB[s][1] == 0 && s+1 < len(AB) {
			s += 1
		}
	}
	fmt.Println(g, " ", score)
}

func D_tool_point_count(s string, x_ret *[]int, p_ret *[]int) {
	for i := 0; i < len(s); i++ {
		//fmt.Println(string(s[i]))
		if string(s[i]) == "X" {
			*x_ret = append(*x_ret, i)
		} else {
			*p_ret = append(*p_ret, i)
		}
	}
}

func D_LongestX(s string, k int) {
	var x_ret []int
	var p_ret []int
	D_tool_point_count(s, &x_ret, &p_ret)
	fmt.Println(x_ret, p_ret)
	for i := 0; i < k; i++ {

	}
}
func main() {
	/*
		fmt.Print("A_Problem\n")
		A_FirstGrid("##", "##")
		A_FirstGrid(".#", ".#")
		A_FirstGrid("#.", ".#")
		fmt.Print("\nB_Problem\n")
		B_HardCalculation("229", "390")
		B_HardCalculation("220", "400")
		B_HardCalculation("329", "380")
		B_HardCalculation("329", "389")
		fmt.Print("\nC_Problem\n")
		a := [][]int{{3, 1}, {4, 2}, {2, 3}}
		//C_Cheese(3, 5, a)
		C_Cheese_New(3, 5, a)
		b := [][]int{{6, 2}, {1, 5}, {3, 9}, {8, 7}}
		C_Cheese_New(4, 100, b)
		c := [][]int{{314944731, 649},
			{140276783, 228},
			{578012421, 809},
			{878510647, 519},
			{925326537, 943},
			{337666726, 611},
			{879137070, 306},
			{87808915, 39},
			{756059990, 244},
			{228622672, 291}}
		C_Cheese_New(10, 3141, c)*/

	fmt.Print("\nD_Problem\n")
	var d1 string = "XX...X.X.X."
	D_LongestX(d1, 2)

}
