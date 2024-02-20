package main

import "fmt"

func main() {
	/*
		terdapat 2 cara penulisan untuk menulis variable. Yang pertama adalah yang
		dituliskan tipedatanya dan ada juga yang tidak ditulis tipe datanya atau biasa.
	*/
	// declaring variable

	// variable with data type
	var name string = "Airell"
	var address string
	var age int = 23

	address = "JL. Raya Wado, Kec. Wado, Kab. Sumedang"

	// variable without data type
	var married = true
	var hight = 165

	// variable without data type (short declaration)
	nama := "Khalinda"
	umur := 16

	// multiple variable declarations
	var student1, student2, student3 string = "satu", "dua", "tiga"
	var first, second, third int
	first, second, third = 1, 2, 4

	var nama1, umur1, status = "danil", 29, true

	var saudara1, saudara2 = "Razil", "Khalinda"

	// underscore variable
	// ketika variabel tidak di pakai
	_, _ = saudara1, saudara2

	fmt.Printf("my biodata : %v, status = %v, umur = %v \n", nama1, status, umur1)
	fmt.Printf("Halo nama saya %s, umurku adalah %d, dan aku tinggal di %s \n", nama, umur, address)
	fmt.Println(student1, student2, student3)
	fmt.Println(first, second, third)

	fmt.Println("Ini adalah namanya =>", name)
	fmt.Println("Ini adalah umurnya =>", age)
	fmt.Println("Ini adalah alamatnya =>", address)
	fmt.Println("Married ? ", married)
	fmt.Println("Tinggi Badan => ", hight)

	// cetak type data variable
	fmt.Printf("%T, %T \n", name, age)

	fmt.Println("Name Sister =>", nama)
	fmt.Println("Umurnya =>", umur)
}
