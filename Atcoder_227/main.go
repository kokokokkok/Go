package main

import (
	"fmt"
)

func A_LastCard(n, k, a int) {
	ret := ((k - a - 1) % n) + 1
	fmt.Println(ret)
}

func B_KEYENCEbuidiing(n int, S []int) {
	ans := 0
	for i := 0; i < len(S); i++ {
		n := false
		for a := 1; a < 1000; a++ {
			for b := 1; b < 1000; b++ {
				mes := 4*a*b + 3*a + 3*b
				if mes == S[i] {
					n = true
				}
			}
		}
		if n == false {
			ans++
		}
	}
	println(ans)
}
func main() {
	fmt.Print("\nA_Problem\n")
	A_LastCard(3, 3, 2)
	A_LastCard(1, 100, 1)
	A_LastCard(3, 14, 2)

	fmt.Print("\nB_Problem\n")
	b1 := []int{10, 20, 39}
	B_KEYENCEbuidiing(3, b1)
	b2 := []int{666, 777, 888, 777, 666}
	B_KEYENCEbuidiing(5, b2)

	fmt.Print("\nC_Problem\n")
	fmt.Print("\nD_Problem\n")
}
