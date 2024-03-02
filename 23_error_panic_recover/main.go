package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		Error merupakan sebuah tipe data pada bahasa Go yang digunakan untuk me-return sebuah error jika ada error yang terjadi. Umumya, nilai error akan di-return pada pada posisi terakhir dari nilai-nilai return sebuah function.
		Ketika tidak ada error yang di return maka variable err akan menjadi nil karena zero value dari tipe data error adalah nil.
	*/
	var number int
	var err error

	number, err = strconv.Atoi("1234DN") // konversi string to int

	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println("terjadi error :", err.Error())
	}

	number, err = strconv.Atoi("123")

	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}
}
