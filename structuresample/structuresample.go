package structuresample

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	age       int
}

type address struct {
	city  string
	state string
}

type person struct {
	firstName string
	lastName  string
	address
}

type rectangle struct {
	length int
	width  int
}

func area(r rectangle) {
	fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}

func (r rectangle) area() {
	fmt.Printf("Area Method result: %d\n", (r.length * r.width))
}

func perimeter(r *rectangle) {
	fmt.Println("perimeter function output:", 2*(r.length+r.width))
}

func (r *rectangle) perimeter() {
	fmt.Println("perimeter method output:", 2*(r.length+r.width))
}

func (a address) fullAddress() {
	fmt.Printf("Full address: %s, %s\n", a.city, a.state)
}

func StructureTest() {
	r := rectangle{
		length: 10,
		width:  5,
	}
	area(r)
	fmt.Println("============================")
	r.area()
	fmt.Println("============================")
	p := &r
	fmt.Println("============================")
	p.area()

	fmt.Println("============================")
	p1 := &r //pointer to r
	perimeter(p1)
	fmt.Println("============================")
	p1.perimeter()
	fmt.Println("============================")
	r.perimeter() //calling pointer receiver with a value
}

func StructureInit() {
	e1 := &Employee{
		lastName:  "Qu",
		firstName: "Kimi",
		age:       36,
	}
	// fmt.Println(e1)
	e2 := &Employee{"Benson", "Wang", 37}
	fmt.Println(e2)
	addr := &e2
	fmt.Println(addr)

	e3 := Employee{
		firstName: "Lucy",
		lastName:  "Zhuang",
		age:       38,
	}
	fmt.Println(e3)

	employee := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Alan",
		lastName:  "Xu",
		age:       34,
	}
	fmt.Println(employee)
	fmt.Println("============================")
	e1.display()
	fmt.Println("============================")
	fmt.Println(e2.firstName)
	e2.changeName("what")
	fmt.Println(e2.firstName)
	fmt.Println("============================")
	fmt.Println(e2.age)
	e2.changeAge(-1)
	fmt.Println(e2.age)
	fmt.Println("============================")
	p := person{
		firstName: "Elon",
		lastName:  "Musk",
		address: address{
			city:  "Los Angeles",
			state: "California",
		},
	}
	p.fullAddress()
}

func (e *Employee) display() {
	fmt.Println(e)
}

func (e Employee) changeName(newName string) {
	e.firstName = newName
}

func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}
