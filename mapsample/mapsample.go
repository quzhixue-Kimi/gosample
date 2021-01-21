package mapsample

import "fmt"

func init() {
	fmt.Println("map test init...")
}

func MapTest() {
	fmt.Println("=================map test==================")
	personSalary := make(map[string]int)
	personSalary["steve"] = 12000
	personSalary["jamie"] = 15000
	personSalary["mike"] = 9000
	fmt.Println("personSalary map contents:", personSalary)
	personSalary["kim"] = 20000
	fmt.Println("personSalary map contents:", personSalary)

	personSalary1 := map[string]int{
		"steve1": 12000,
		"jamie1": 15000,
	}
	personSalary1["mike1"] = 9000
	fmt.Println("personSalary1 map contents:", personSalary1)
	fmt.Println("=================map test==================")
	if value, ok := personSalary["kim"]; ok {
		fmt.Println("salary is ", value)
	}
	fmt.Println("=================map test==================")
	for key, value := range personSalary {
		fmt.Printf("personSalary[%s] = %d\n", key, value)
	}
	fmt.Println("length is", len(personSalary))
	fmt.Println("=================map test==================")
	delete(personSalary, "kim")
	for key, value := range personSalary {
		fmt.Printf("personSalary[%s] = %d\n", key, value)
	}
	fmt.Println("length is", len(personSalary))
	fmt.Println("=================map test==================")
}
