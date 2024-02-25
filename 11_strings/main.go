package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	city := "Jakarta"

	// string merupakan kumpulan byte angka decimal ASCII code
	for i := 0; i < len(city); i++ {
		fmt.Printf("Index: %d, byte: %d\n", i, city[i])
	}
	fmt.Printf("byte J : %d , stringnya: %v ", city[0], string(city[0]))

	fmt.Println(strings.Repeat("#", 50))

	// membuat string menggunakan slice of byte
	var nama []byte = []byte{100, 97, 110, 105, 108}

	fmt.Printf("nama saya %v", string(nama))

	fmt.Println(strings.Repeat("#", 50))

	// rune-by-rune
	str1 := "makan"
	str2 := "makân" //mengandung accented-character

	// mencari jumlah byte pada string
	fmt.Printf("str1 byte length = %d\n", len(str1))
	fmt.Printf("str2 byte length = %d\n", len(str2))

	// mencari jumlah karakter pada string
	fmt.Printf("str1 character length => %d\n", utf8.RuneCountInString(str1))
	fmt.Printf("str2 character length => %d\n", utf8.RuneCountInString(str2))

	// looping string secara rune-per-rune menggunakan range loop
	str3 := "makân"

	for i, s := range str3 {
		fmt.Printf("index => %d, string => %s\n", i, string(s))
	}

	// noted : ketika ada karakter yang tidak mendukung ASCII code maka di hitung lebih dari 1 index
}
