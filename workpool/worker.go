package workpool

import (
	"fmt"
	"time"
)

func WorkerPool(id int, jobs chan int, results chan int) {
	for i := range jobs {
		fmt.Printf("worker %d starts job %d\n", id, i)
		time.Sleep(time.Millisecond * 900)
		fmt.Printf("worker %d ends job %d\n", id, i)
		results <- i * i
	}
	//close(results)
}
