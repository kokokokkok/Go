package main

import (
	"fmt"
	"time"
)

func say(s string) { //しばらく待ってからログ出力を行う
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

/*
func finonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}*/

func finonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x: //cに送信された場合とquitに送信された場合で切り替える
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func main() {
	/*
		go say("world")
		say("hello")
		s := []int{7, 2, 8, -9, 4, 0}
		c := make(chan int)
		go sum(s[:len(s)/2], c)//cのチャネルが非同期で並列で行われている
		go sum(s[len(s)/2:], c)
		x, y := <-c, <-c //ｃが準備でき来た順でx,yに入れる
		fmt.Println(x, y, x+y)

		ch := make(chan int, 2)
		ch <- 1
		ch <- 2
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		//fmt.Println(<-ch)
		c := make(chan int, 10)
		go finonacci(cap(c), c)
		for i := range c { //　close(c)でチャネルを閉じなければ、つかされ続ける配列を永続的に出力し続ける事になる
			fmt.Println(i)*/
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	finonacci(c, quit)

}
