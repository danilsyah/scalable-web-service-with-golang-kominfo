package main

import "fmt"

func main() {
	/*
		Constant (const) atau Konstanta merupakan jenis variable pada bahasa Go yang nilainya tidak dapat diubah
		ketika kita membuat variable denganconst, maka kita harus langsung memberikan nilai kepadanya.
		Karena jika tidak maka akan timbul error pada saat compile time

		1. Operator Aritmatika
			+ , - , * , / , %, ++ , --
			level prioritas perhitungan
			kukabataku = (), *, /, +, -

		2. Operator relational atau perbandingan
			==, !=, >, <, >=, <=

		3. Operator Logika
			&& , ||, !
	*/

	const full_name string = "Haykal Dafiansyah"

	// aritmatika
	var value = 2 + 2*3 // result 8
	// gunakan tanda kurung () jika ingin penghitungan di awal
	var value2 = (2 + 2) * 3 // result 12

	// operator perbandingan
	var firstCondition bool = 2 < 3
	var secondCondition bool = "danil" == "Danil"
	var thirdCondition bool = 10 != 2.3
	var fourthCondition bool = 11 <= 11

	// operator logika
	var right = true
	var wrong = false

	var wrongAndRight = wrong && right
	fmt.Printf("wrong && right \t(%t) \n", wrongAndRight)

	var wrongOrRight = wrong || right
	fmt.Printf("wrong || right \t(%t) \n", wrongOrRight)

	var wrongReverse = !wrong
	fmt.Printf("!wrong \t\t(%t) \n", wrongReverse)

	fmt.Printf("Hello %s\n", full_name)
	fmt.Println("result = ", value)
	fmt.Println("result2 = ", value2)

	fmt.Println("first condition: ", firstCondition)
	fmt.Println("second condition: ", secondCondition)
	fmt.Println("third condition: ", thirdCondition)
	fmt.Println("fourth condition: ", fourthCondition)
}
