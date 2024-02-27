package main

import (
	"fmt"
	"strings"
)

// alias callback
type isOddNum func(int) bool

func main() {

	/*
		closure merupakan sebuah anonymous function yang dapat disimpan sebagai sebuah variable
		maupun dapat dijadikan sebagai nilai return dari sebuah function
	*/

	// declare closure in variabel
	var evenNumbers = func(numbers ...int) []int {
		var result []int

		for _, v := range numbers {
			if v%2 == 0 {
				result = append(result, v)
			}
		}

		return result
	}

	var numbers = []int{4, 93, 77, 10, 52, 22, 34}

	fmt.Println(evenNumbers(numbers...))

	// clouse IIFE (immediately-invoked function expression)
	var isPalindrome = func(str string) bool {
		var temp string = ""

		for i := len(str) - 1; i >= 0; i-- {
			temp += string(byte(str[i]))
		}

		return temp == str
	}("katak")

	fmt.Println(isPalindrome)

	fmt.Println(strings.Repeat("#", 50))

	var studentLists = []string{"Danil", "Haykal", "Nufika", "Razil"}

	var find = findStudent(studentLists)

	fmt.Println(find("razil"))

	// closure callback
	/*
		callback adalah sebuah closure yang dijadikan sebagai parameter pada sebuah function.
	*/
	var numbers2 = []int{2, 5, 8, 10, 3, 99, 23}

	var find2 = findOddNumbers(numbers2, func(number int) bool {
		return number%2 != 0
	})

	fmt.Println("Total odd numbers: ", find2)

	fmt.Println(strings.Repeat("*", 50))

	var find3 = findOddNumbers2(numbers2, func(number int) bool {
		return number%2 != 0
	})
	fmt.Println("Total odd numbers:", find3)
}

// closure as a return value
func findStudent(students []string) func(string) string {
	return func(s string) string {
		var student string
		var position int

		for i, v := range students {
			if strings.ToLower(v) == strings.ToLower(s) {
				student = v
				position = i
				break
			}
		}
		if student == "" {
			return fmt.Sprintf("%s does'nt exist!!!", s)
		}
		return fmt.Sprintf("We found %s at position %d", s, position)
	}
}

// closure callback
func findOddNumbers(numbers []int, callback func(int) bool) int {
	var totalOddNumbers int

	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}
	return totalOddNumbers
}

// closure callback alias
func findOddNumbers2(numbers []int, callback isOddNum) int {
	var totalOddNumbers int

	for _, v := range numbers {
		if callback(v) {
			totalOddNumbers++
		}
	}
	return totalOddNumbers
}
