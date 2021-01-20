package main

import (
	"fmt"
	"os"
)

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const (
	n1 = 100
	n2
	n3
)

func GetData() (int, int) {
	return 100, 200
}
func main() {
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)

	fmt.Println("this is a test line")
	fmt.Println("what is up?")
	fmt.Printf("helloworld\n")
	fmt.Println(os.Args)

	if len(os.Args) > 1 {
		fmt.Println("The parameter is", os.Args[1])
	}

	a, _ := GetData()
	_, b := GetData()
	fmt.Println(a, b)

	//声明局部变量 a 和 b 并赋值
	var a1 int = 3
	var b1 int = 4
	//声明局部变量 c 并计算 a 和 b 的和
	c := a1 + b1
	fmt.Printf("a1 = %d, b1 = %d, c = %d\n", a1, b1, c)

	d := sum(a1, b1)
	fmt.Printf("a1 = %d, b1 = %d, c = %d\n", a1, b1, d)

	fmt.Printf("result = %d\n", btoi(false))

	var str = "C语言中文网\nGo语言教程"
	fmt.Println(str)

	s := "hel" + "lo,"
	s += "world!"
	fmt.Println(s) //输出 “hello, world!”
	fmt.Println(Monday)
	os.Exit(0)
}

func sum(a, b int) int {
	fmt.Printf("sum() 函数中 a = %d\n", a)
	fmt.Printf("sum() 函数中 b = %d\n", b)
	num := a + b
	return num
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
