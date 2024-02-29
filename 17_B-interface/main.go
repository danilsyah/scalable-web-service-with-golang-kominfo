package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		Empty interface merupakan suatu tipe data yang dapat menerima segala tipe data pada bahasa Go. Maka dari itu, sebuah variable dengan tipe data empty interface dapat diberikan nilai dengan tipe data yang berbeda-beda.
	*/

	// deklarasi variabel empty interface
	var randomValues interface{}

	_ = randomValues

	randomValues = "Jalan Sumedang"

	randomValues = 1994

	randomValues = true

	randomValues = []string{"Danil", "Syah"}

	fmt.Println(strings.Repeat("#", 50))

	// empty interface type assertion
	var v interface{}
	v = 20

	// agar variabel empty interface dapat melakukan operasi aritmatika
	// maka harus melakukan konversi ke type assertion
	if value, ok := v.(int); ok == true {
		v = value * 9
	}

	// Slice with Empty Interface
	rs := []interface{}{1, "danil", true, 2, "haykal", true}

	rm := map[string]interface{}{
		"Name":   "Danil",
		"Status": true,
		"Age":    29,
	}

	fmt.Printf("%v\n", rs)
	fmt.Printf("%v", rm)
}
