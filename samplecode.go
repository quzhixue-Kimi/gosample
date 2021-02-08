package main

import "github.com/quzhixue-Kimi/gosample/mysqlsample"

//func sendData(ch1 chan int) {
//	for i := 0; i < 10; i++ {
//		ch1 <- i
//		time.Sleep(500 * time.Millisecond)
//		fmt.Println("写入数据 ", i)
//	}
//	close(ch1)
//}
//
//func readData(ch1 chan int, exit chan bool) {
//	for {
//		v, ok := <-ch1
//		if !ok {
//			fmt.Println("已经读取了所有的数据，", ok)
//			break
//		}
//		time.Sleep(700 * time.Millisecond)
//		fmt.Println("取出数据：", v)
//	}
//	exit <- true
//	close(exit)
//}

func main() {
	//	mysqlsample.MySqlSample()
	mysqlsample.MySqlSample1()
	//jobs := make(chan int, 5)
	//results := make(chan int, 5)

	//for i := 1; i <= 2; i++ {
	//	go workpool.WorkerPool(i, jobs, results)
	//}

	//for i := 1; i <= 5; i++ {
	//	jobs <- i
	//}
	//close(jobs)

	////for v := range results {
	////	fmt.Println(v)
	////}

	////for {
	////	if v, ok := <-results; !ok {
	////		break
	////	} else {
	////		fmt.Println(v)
	////	}
	////}

	//for i := 1; i <= 5; i++ {
	//	x := <-results
	//	fmt.Println(x)
	//}

	//	selectsample.SelectChannel()

	//msg := make(chan *io.LogInfo, 5)
	//exit := make(chan bool, 1)

	//l := io.NewLogInfo("hello", "Info")

	//for {
	//	go l.FileReadOperator(msg)
	//	go l.FileCreateOperator(msg, exit)
	//	if _, ok := <-exit; !ok {
	//		break
	//	}
	//}

	//ch1 := make(chan int, 5)
	//exit := make(chan bool, 1)
	//for {
	//	go sendData(ch1)
	//	go readData(ch1, exit)
	//	if _, ok := <-exit; !ok {
	//		break
	//	}
	//}
}
