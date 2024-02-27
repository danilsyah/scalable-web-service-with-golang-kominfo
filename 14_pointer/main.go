package main

import (
	"fmt"
	"strings"
)

/*
	pointer merupakan sebuah variabel spesial yang digunakan untuk menyimpan alamat memory variable
	sebagai contoh sebuah variable bertipe integer memiliki nilai 4, maka yang dimaksud pointer adalah
	alamat memori dimana nilai 4 disimpan, bukan 4 itu sendiri.
	variabel-variabel yang memiliki reference atau alamat memori yang sama, saling berhubungan satu sama lain
	dan nilainya pasti sama, ketika ada perubahan nilai, maka akan memberikan efek kepada variabel lain
	(alamat memorinya sama) yaitu nilainya juga ikut berubah.
*/

func main() {
	// variabel value
	var firstNumber int = 4

	// variabel pointer
	var secondNumber *int = &firstNumber

	fmt.Println("firstNumber (value)	:", firstNumber)
	fmt.Println("firstNumber (memori address) :", &firstNumber)

	fmt.Println("secondNumber (value)	:", *secondNumber)
	fmt.Println("secondNumber (memori address) :", secondNumber)

	fmt.Println(strings.Repeat("*", 50))

	var a int = 10
	fmt.Println("Before:", a)

	changeValue(&a)

	fmt.Println("After:", a)

	var name string = "danil"
	fmt.Println("Before name:", name)
	changeName(&name)
	fmt.Println("After name:", name)

}

// pointer as a paremeter
func changeValue(number *int) {
	*number = 20
}

func changeName(name *string) {
	*name = "Haykal"
}
