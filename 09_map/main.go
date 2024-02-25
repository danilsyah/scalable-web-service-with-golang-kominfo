package main

import (
	"fmt"
	"strings"
)

func main() {
	// map salah satu tipe data reference type
	// zero value dari tipe data map adalah nil

	var person map[string]string // Deklarasi

	person = map[string]string{} // Inisialisasi

	person["name"] = "Danil"
	person["age"] = "29"
	person["address"] = "Jalan Wadogirang"

	fmt.Println("name: ", person["name"])
	fmt.Println("age: ", person["age"])
	fmt.Println("address: ", person["address"])

	fmt.Println(strings.Repeat("#", 50))

	// memberikan key dan value langsung ketika deklarasi variable map
	var car = map[string]string{
		"merk":  "honda",
		"warna": "hitam",
		"jenis": "sedan",
	}

	for key, value := range car {
		fmt.Println(key, ":", value)
	}

	fmt.Println("Before deleting: ", car)

	// delete value map
	delete(car, "warna")

	fmt.Println("After deleting:", car)

	fmt.Println(strings.Repeat("#", 50))

	var fruit = map[string]string{
		"nama":  "mangga",
		"berat": "2kg",
		"warna": "orange",
	}

	// Detecting a value
	value, exist := fruit["berat"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("value does'nt exist ")
	}

	delete(fruit, "berat")

	value, exist = fruit["berat"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("value has been deleted")
	}

	// combining slice with map
	var people = []map[string]string{
		{"name": "danil", "age": "29"},
		{"name": "haykal", "age": "5"},
		{"name": "fika", "age": "29"},
	}

	for i, person := range people {
		fmt.Printf("Index: %d, name: %s, age: %s\n", i, person["name"], person["age"])
	}

}
