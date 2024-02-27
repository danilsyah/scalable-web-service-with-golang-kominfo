package main

import (
	"fmt"
	"strings"
)

type Employee struct {
	name     string
	age      int
	division string
}

type Person struct {
	name string
	age  int
}

type Animal struct {
	name string
	age  int
}

type Food struct {
	makanan string
	animal  Animal // Embedded struct
}

func main() {
	var employee Employee

	employee.name = "Danil"
	employee.age = 29
	employee.division = "Technology Information"

	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.division)

	fmt.Println(strings.Repeat("*", 50))

	var employee1 = Employee{}
	employee1.name = "Haykal"
	employee1.age = 6
	employee1.division = "Devops"

	// initializing struct sekaligus memberikan nilai-nilai
	var employee2 = Employee{name: "Nufika", age: 28, division: "Finance"}

	fmt.Printf("Employee1: %+v\n", employee1)
	fmt.Printf("Employee2: %+v\n", employee2)

	var employee3 = Employee{name: "Khalinda", age: 30, division: "Marketing"}

	// pointer to a struct
	var employee4 *Employee = &employee3

	fmt.Println("Employee3 name : ", employee3.name)
	fmt.Println("Employee4 name : ", employee4.name)

	fmt.Println(strings.Repeat("=", 50))

	employee4.name = "Razil"

	fmt.Println("Employee3 name : ", employee3.name)
	fmt.Println("Employee4 name : ", employee4.name)

	var food = Food{}

	food.animal.name = "Harimau"
	food.animal.age = 3
	food.makanan = "karnivora"

	fmt.Printf("%+v", food)

	fmt.Println(strings.Repeat("*", 50))

	// Anonymous struct tanpa pengisian field
	var employee5 = struct {
		person   Person
		division string
	}{}
	employee5.person = Person{name: "Dikri", age: 24}
	employee5.division = "management"

	// anonymous struct dengan pengisian field
	var employee6 = struct {
		person   Person
		division string
	}{
		person:   Person{name: "Haykal", age: 5},
		division: "Owner",
	}

	fmt.Printf("Employee5: %+v\n", employee5)
	fmt.Printf("Employee6: %+v\n", employee6)

	fmt.Println(strings.Repeat("*", 50))

	// slice of struct
	var people = []Person{
		{name: "Danil", age: 29},
		{name: "Nufika", age: 28},
		{name: "Haykal", age: 5},
	}

	for _, v := range people {
		fmt.Printf("%+v\n", v)
	}

	fmt.Println(strings.Repeat("*", 50))

	// slice of anonymous struct
	var employee7 = []struct {
		person   Person
		division string
	}{
		{person: Person{name: "danil", age: 29}, division: "IT Infrastucture"},
		{person: Person{name: "haykal", age: 5}, division: "Owner"},
		{person: Person{name: "nufika", age: 29}, division: "Finance"},
	}

	for _, v := range employee7 {
		fmt.Printf("%+v\n", v)
	}
}
