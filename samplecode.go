package main

import (
	"fmt"
	"time"
)

func sendData(ch1 chan int) {
	for i := 0; i < 10; i++ {
		ch1 <- i
		time.Sleep(500 * time.Millisecond)
		fmt.Println("写入数据 ", i)
	}
	close(ch1)
}

func readData(ch1 chan int, exit chan bool) {
	for {
		v, ok := <-ch1
		if !ok {
			fmt.Println("已经读取了所有的数据，", ok)
			break
		}
		time.Sleep(700 * time.Millisecond)
		fmt.Println("取出数据：", v)
	}
	exit <- true
	close(exit)
}

func main() {
	ch1 := make(chan int, 5)
	exit := make(chan bool, 1)
	for {
		go sendData(ch1)
		go readData(ch1, exit)
		if _, ok := <-exit; !ok {
			break
		}
	}
}
