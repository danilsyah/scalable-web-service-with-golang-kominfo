package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	greet("Danil", 29)
	intro("haykal", "wadogirang")

	var names = []string{"danil", "syah"}
	var printMsg = greet1("Hello", names)

	fmt.Println(printMsg)

	var diameter float64 = 15

	// menghitung luas dan keliling lingkaran
	var area, circumference = calculate(diameter)

	fmt.Println("Area:", area)
	fmt.Println("Circumference:", circumference)

	var area1, circumference1 = calculate2(diameter)

	fmt.Println("Area:", area1)
	fmt.Println("Circumference:", circumference1)

	// variadic function
	studentLists := print("Danil", "Fika", "Haykal", "Razil", "Dikri", "Khalinda")

	fmt.Printf("%v\n", studentLists)

	// variadic function sum
	numberLists := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := sum(numberLists...)
	fmt.Println("Result:", result)

	profile("danil", "pasta", "ayah geprek", "ikan loa", "sate padang")
}

func greet(name string, age int8) {
	fmt.Printf("Hello there!, My name is %s and I'm %d years old", name, age)
}

// func ketika tipe data parameter sama
func intro(name, address string) {
	fmt.Println("Hello there! My name is", name)
	fmt.Println("I live in ", address)
}

// function return
func greet1(msg string, names []string) string {
	var joinStr = strings.Join(names, " ")

	var result string = fmt.Sprintf("%s %s", msg, joinStr)

	return result
}

// returning multiple values
func calculate(d float64) (float64, float64) {
	// menghitung luas
	var area float64 = math.Pi * math.Pow(d/2, 2)

	// menghitung keliling
	var circumference = math.Pi * d

	return area, circumference
}

// function predefined return value
// membuat sebuat variabel sebagai hasil nilai return
func calculate2(d float64) (area float64, circumference float64) {
	// menghitung luas
	area = math.Pi * math.Pow(d/2, 2)

	// menghitung keliling
	circumference = math.Pi * d

	return
}

// function dengan parameter tidak terbatas
// variadic function
func print(names ...string) []map[string]string {
	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("student%d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}

	return result
}

// variadic function2
func sum(numbers ...int) int {
	total := 0

	for _, v := range numbers {
		total += v
	}
	return total
}

func profile(name string, favFoods ...string) {
	mergeFavFoods := strings.Join(favFoods, ",")

	fmt.Println("Hello there!!! I'm", name)
	fmt.Println("I really love to eat", mergeFavFoods)
}
