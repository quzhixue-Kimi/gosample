package main

import (
	"fmt"
	"math"
	"os"
	"unsafe"

	abc "github.com/quzhixue-Kimi/gosample/mapsample"
	point "github.com/quzhixue-Kimi/gosample/pointersample"
	sli "github.com/quzhixue-Kimi/gosample/slicesample"
	str "github.com/quzhixue-Kimi/gosample/stringsample"
	s "github.com/quzhixue-Kimi/gosample/structuresample"
	"github.com/quzhixue-Kimi/stringutil"
)

// type Weekday int

const (
	Sunday    = iota // 0
	Monday    = 100  // 100
	Tuesday          // 100
	Wednesday = 200  // 200
	Thursday  = iota // 4
	_                // 5
	Friday           // 6
	Saturday         // 7
)

const typedString string = "how old are you?"

const (
	n1 = 100
	n2
	n3
)

func forTest() {
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d", i)
	}
}

func forTest1() {
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break //loop is terminated if i > 5
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nline after for loop")
}

func condTest() {
	if num := 10; num%2 == 0 {
		fmt.Println("the number is even")
	} else {
		fmt.Println("the number is odd")
	}
}

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

func forTest2() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
}

func number() int {
	num := 15 * 5
	return num
}

func switchTest() {

	switch finger := 8; finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default:
		fmt.Println("incorrect finger number")
	}

	letter := "i"
	switch letter {
	case "a", "e", "i", "o", "u": //multiple expressions in case
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}

	num := 75
	switch { // expression is omitted
	case num >= 0 && num <= 50:
		fmt.Println("num is greater than 0 and less than 50")
	case num >= 51 && num <= 100:
		fmt.Println("num is greater than 51 and less than 100")
	case num >= 101:
		fmt.Println("num is greater than 100")
	}

	switch num1 := number(); {
	case num1 < 50:
		fmt.Printf("%d is lesser than 50\n", num1)
		fallthrough
	case num1 < 100:
		fmt.Printf("%d is lesser than 100\n", num1)
		fallthrough
	case num1 < 200:
		fmt.Printf("%d is lesser than 200\n", num1)
	}
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

func arrayTest() {
	var a [3]int
	a[0] = 12 // array index starts at 0
	a[1] = 78
	a[2] = 50
	fmt.Println(a)

	b := [...]int{110, 21, 222, 42, 10}
	fmt.Println(b)

	str1 := [...]string{"USA", "China", "India", "Germany", "France"}
	fmt.Println(str1)
	str2 := str1
	str2[0] = "Mexco"
	fmt.Println(str1)
	fmt.Println(str2)

	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function ", num)
	changeLocal(num) //num is passed by value
	fmt.Println("after passing to function ", num)
	fmt.Println(len(num))
	var sum = 0
	for i, v := range num {
		fmt.Printf("%d the element of a is %v\n", i, v)
		sum += v
	}
	fmt.Println(sum)
}

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)
}

func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func multiArray() {
	a := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
	}
	printarray(a)
	var b [3][2]string
	b[0][0] = "apple"
	b[0][1] = "samsung"
	b[1][0] = "microsoft"
	b[1][1] = "google"
	b[2][0] = "AT&T"
	b[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printarray(b)
}

func sliceTest() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4] //creates a slice from a[1] to a[3]
	fmt.Println(b)
	c := []int{6, 7, 8} //creates and array and returns a slice reference
	fmt.Println(c)
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5] //90,82,100
	fmt.Println(dslice)
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println(darr)
	numa := [3]int{78, 79, 80}
	nums1 := numa[:] //creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array before change 1", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice))
	fmt.Println(fruitslice)
	fruitslice = fruitslice[:cap(fruitslice)]
	fmt.Println("After re-slicing length is", len(fruitslice), "and capacity is", cap(fruitslice))
}

func sliceMakeTest() {
	slice := make([]int, 4, 5)
	fmt.Println(slice)
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars))
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars))
	var names []string
	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "John", "Sebastian", "Vinay")
		fmt.Println("names contents:", names)
	}
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:", food)
	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	subtactOne(nos)
	fmt.Println("slice after function call", nos)
	pls := [][]string{
		{"C", "C++"},
		{"JavaScript", "java"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func sliceGc() {
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)
}

func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy

	countries[0] = "China"
	fmt.Println(countries)
	return countriesCpy
}

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func findIndex() {
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	find(87)

	fmt.Println("=================================")
	n := []int{10, 23, 41}
	find(41, n...)
}

func main() {
	sli.SliceCopy()
	fmt.Println("================================")
	sli.SliceChange()
	fmt.Println("=======================")
	sli.SliceScale()
	fmt.Println("======================")
	sli.SliceTest2()
	fmt.Println("================================")
	sli.SliceTest1()
	fmt.Println("================================")
	sli.SliceTest()
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)

	s.StructureInit()
	s.StructureTest()
	point.PointerTest()
	str.PrintBytes("hello world")
	str.PrintChars("hello world")
	str.PrintChars("Señor")
	str.PrintChars1("Señor")
	str.PrintString("Señor")
	str.Length("Señor")
	fmt.Println(str.Mutate([]rune("hello")))
	// mapsample.MapTest()
	abc.MapTest()
	findIndex()
	sliceGc()
	sliceMakeTest()
	sliceTest()
	multiArray()
	arrayTest()
	switchTest()
	forTest()
	fmt.Println()
	forTest1()
	fmt.Println()
	forTest2()
	fmt.Println()
	condTest()
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
	age = 21
	fmt.Println("my age is :", age)

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
	sp := fmt.Sprintf("Hi man, %s %s %s", "what", "is", "up?")
	fmt.Println(sp)
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
