package main

import "fmt"

func main() {
	// array of int
	var numbers [4]int

	numbers = [4]int{1, 2, 3, 4}

	// array of string
	var names = [3]string{"danil", "fika", "haykal"}

	fmt.Printf("%#v\n", numbers)
	fmt.Printf("%#v\n", names)

	// modify element through index / ubah elemen melalui index
	var fruits = [3]string{"apel", "pisang", "mangga"}
	fruits[0] = "apple"
	fruits[1] = "banana"
	fruits[2] = "mango"

	fmt.Printf("%#v\n", fruits)

	// mengakses elemen array

	// cara pertama
	for i, v := range fruits {
		fmt.Printf("Index: %d, value: %s\n", i, v)
	}

	// cara kedua
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("index - %v, value : %v\n", i, fruits[i])
	}

	// array multidimensional
	/*
		Tanda [2][3]int  memiliki arti bahwa array terluar nya memiliki panjang sama dengan 2,
		dan array yang berada di dalamnya memiliki panjang sama dengan 3 dengan tipe data int
	*/
	balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}

	for i, arr := range balances {
		for _, value := range arr {
			fmt.Printf("index %d = %d ", i, value)
		}
		fmt.Printf("\n")
	}
}
