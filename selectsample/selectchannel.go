package selectsample

import (
	"fmt"
	"time"
)

func SelectChannel() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(200 * time.Millisecond)
		for i := 1; i <= 10; i++ {
			ch1 <- i
		}

	}()
	go func() {
		time.Sleep(200 * time.Millisecond)
		for i := 100; i <= 110; i++ {
			ch2 <- i
		}
		//close(ch2)
	}()

	for {
		time.Sleep(time.Second * 2)
		select {
		case num1 := <-ch1:
			fmt.Println("ch1中取数据。。", num1)
		case num2, ok := <-ch2:
			if ok {
				fmt.Println("ch2中取数据。。", num2)
			} else {
				fmt.Println("ch2通道已经关闭。。")
			}
		}
	}
}
