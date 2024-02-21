package main

import "fmt"

func main() {
	/*
		Tipe data pada bahasa Go terbagi menjadi 4 kategori dengan detail seperti berikut ini:
		1.Basic Type: number, string, boolean.
		2.Aggregate Type: array dan struct.
		3.Reference Type: slice, pointer, map, function, channel
		4.Interface Type: interface

		- tipe data integer non desimal
		uint8 		= 0 - 255
		uint16 		= 0 - 65535
		uint32		= 0 - 4294967295
		uint64		= 0 - 18446744079551615
		int8		= -128 - 127
		int16		= -32768 - 32767
		int32		= -2147483648 - 2147483647
		int64		= -9223372036854775808 - 9223372036854775807

		- data type numerik desimal pecahan
		float32
		float64

		- tipe data bool
		bool = true | false

		- tipe data string
		string = nilai di apit oleh  "" (petik dua) atau `` (backticks)

		- nil = nil bukan merupakan tipe data, melainkan sebuah nilai yang benar-benar kosong.
		Variabel yang isi nilainya nil berarti memiliki nilai kosong. tipe data yang di set nil :
			- pointer
			- tipe data function
			- slice
			- map
			- channel
			- interface kosong atau interface{}

		- zero value = nilai default tipe data ketika variable belum diberikan sebuah nilai
			- Zero value dari string adalah "" (string kosong).
			- Zero value dari bool adalah false.
			- Zero value dari tipe numerik non-desimal adalah 0.
			- Zero value dari tipe numerik desimal adalah 0.0
	*/
	var first uint8 = 89
	var second int8 = -12

	var decimalNumber float32 = 3.14

	var isMarried bool = true

	var message string = "Hello word\n"
	var pesan string = `hai nama saya "danil syah" :), 
		mari kita belajar "Scalable Web Service With Go`

	fmt.Printf("tipe data first : %T\n ", first)
	fmt.Printf("tipe data second : %T\n", second)

	fmt.Printf("Decimal number : %f\n", decimalNumber)   // 3.140000
	fmt.Printf("Decimal number : %.2f\n", decimalNumber) // 3.14

	fmt.Printf("Apakah sudah menikah : %t\n", isMarried)

	fmt.Print(message)
	fmt.Print(pesan)
}
