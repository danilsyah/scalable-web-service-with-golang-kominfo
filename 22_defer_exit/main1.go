package main

import "fmt"

func main() {
	callDeferFunc()
	fmt.Println("Hai everyone!!")
}

func callDeferFunc() {
	defer deferFunc()
	// fmt.Println("call defer func")
}

func deferFunc() {
	fmt.Println("Defer func starts to execute")
}
