package main

import "fmt"

/*
	Defer merupakan sebuah keyword pada bahasa Go yang digunakan untuk memanggil sebuah function yang dimana proses eksekusi akan di tahan hingga block sebuah function selesai
*/
func main() {
	defer fmt.Println("defer function starts to execute")
	fmt.Println("Hai everyone")
	fmt.Println("Welcome back to Go learning center")
}
