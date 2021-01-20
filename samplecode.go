package main

import (
	"fmt"
	"math"
	"os"
	"unsafe"

	"github.com/quzhixue-Kimi/stringutil"
)

// type Weekday int

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

const typedString string = "how old are you?"

const (
	n1 = 100
	n2
	n3
)

func GetData() (int, int) {
	return 100, 200
}

func typeTest() {
	const a = 5
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)

	var i = 5
	var f = 5.6
	var c = 5 + 6i
	fmt.Printf("i's type %T, f's type %T, c's type %T", i, f, c)
	fmt.Println()

	var a1 = 5.9 / 8
	fmt.Printf("a1's type %T value %v", a1, a1)
	fmt.Println()
}

func main() {

	typeTest()
	fmt.Printf("%T, %v\n", typedString, typedString)
	fmt.Printf("%T, %v\n", n1, n1)
	var (
		name   = "Kimi"
		age    = 36
		gender = "male"
	)

	fmt.Println("name is :", name, ", age is :", age, ", gender is :", gender)

	//	var age int
	fmt.Println("my age is :", age)
	age = 10
	fmt.Println("my age is :", age)
	fmt.Println(Sunday)
	age = 21
	fmt.Println("my age is :", age)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)
	var width, height int = 100, 50
	a2, b2 := 102, 103
	fmt.Println(a2, b2)
	b2, c2 := 104, 105
	fmt.Println(b2, c2)
	b2, c2 = 106, 107
	fmt.Println(b2, c2)
	a3, b3 := 12.33, 10.22
	c3 := math.Min(a3, b3)
	fmt.Println(c3)
	fmt.Println("width is :", width, ", height is :", height)
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)

	name1, age1 := "Benson", 36
	fmt.Println("name1 ", name1, ", age1 ", age1)
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	fmt.Println(os.Args)

	fmt.Printf("type %T value %v\n", name, name)
	if len(os.Args) > 1 {
		fmt.Println("The parameter is", os.Args[1])
	}

	a, _ := GetData()
	_, b := GetData()
	fmt.Println(a, b)
	var a1 int = 3
	var b1 int = 4
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
	boolTest()
	intTest()
	floatTest()
	stringTest()
	plusTest1()
	area, perimeter := rectProps(2, 3)
	fmt.Printf("Area %f Perimeter %f", area, perimeter)
	os.Exit(0)
}

func plusTest1() {
	i := 55
	j := 67.23
	sum := i + int(j)
	var k float64 = float64(i)
	fmt.Println(sum)
	fmt.Println(k)
}

func stringTest() {
	first := "kimi"
	last := "qu"
	name := first + "." + last
	fmt.Println("my name is :", name)
}

func floatTest() {
	fmt.Println()
	a, b := 5.67, 8.97
	fmt.Printf("type of a %T b %T\n", a, b)
	sum := a + b
	diff := a - b
	fmt.Println("sum", sum, "diff", diff)

	no1, no2 := 56, 89
	fmt.Println("sum", no1+no2, "diff", no1-no2)
}

func intTest() {
	var a int = 89
	b := 95
	fmt.Println("value of a is", a, "and b is", b)
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a))
	fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b))
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

func boolTest() {
	a := true
	b := false
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a && b)
	d := a || b
	fmt.Println(d)
}

func rectProps(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return
}
