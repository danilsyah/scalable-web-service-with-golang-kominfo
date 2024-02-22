package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
			Slice termasuk dalam kategori reference type yang dimana jika kita melakukan copy terhadap suatu slice,
		dan kita ubah element dari yang kitacopytersebut, maka slice semulanya juga akan ikut terubah.
			Slice tidak memiliki sifat fixed-length  yang berarti panjang dari slice tidak tetap sehingga
		kita bisa dengan leluasa menentukan panjang dari slice nya
	*/

	// slice of string
	var fruits = []string{"apple", "banana", "mango"}
	_ = fruits

	fmt.Printf("%#v\n", fruits)

	// new declaration slice of string
	var names = make([]string, 3)
	_ = fruits

	// menambahkan element pada slice kosong
	// cara pertama
	names[0] = "danil"
	names[1] = "fika"
	names[2] = "haykal"

	// cara kedua
	names = append(names, "razil", "khalinda")
	fmt.Printf("%#v\n", names)

	var fruits1 = []string{"apple", "banana", "mango"}
	var fruits2 = []string{"durian", "strawbery", "cerry"}

	// menambahkan element pada slice dari slice lain menggunakan tanda ellipsis (...)
	fruits1 = append(fruits1, fruits2...)
	fmt.Printf("%#v\n", fruits1)

	// copy slice
	var fruits3 = []string{"kelapa", "singkong", "pepaya"}
	var fruits4 = []string{"rambutan", "duren", "nangka"}

	// semua element slice fruits3 akan di replace oleh slice fruits4
	nn := copy(fruits3, fruits4)

	fmt.Println("Fruits3 => ", fruits3)
	fmt.Println("Fruits4 => ", fruits4)
	fmt.Println("Copied elements", nn)

	var sayuran = []string{"wortel", "tomat", "kangkung", "brokoli", "jagung"}

	// slice slicing [start:stop]
	var sayuran1 = sayuran[1:4]
	fmt.Printf("%#v\n", sayuran1)

	var sayuran2 = sayuran[0:]
	fmt.Printf("%#v\n", sayuran2)

	var sayuran3 = sayuran[:3]
	fmt.Printf("%#v\n", sayuran3)

	var sayuran4 = sayuran[:]
	fmt.Printf("%#v\n", sayuran4)

	// combining slicing and append
	// me-replace index ke-3 seterusnya dengan element bayam
	sayuran = append(sayuran[:3], "bayam")
	fmt.Printf("%#v\n", sayuran)
	fmt.Println("Panjang slice sayuran => ", len(sayuran))
	fmt.Println("Kapasitas slice sayuran => ", cap(sayuran))

	// backing array
	// slice akan me reference by memory pada setiap element nya
	// jika melakukan replace element pada cities2 maka akan berubah juga pada element cities
	// Ini terjadi karena variable cities dan cities2 masih dalam satu backing array yang sama (memory)
	// slice lebih hemat memory dibanding array
	var cities = []string{"sumedang", "bandung", "bogor", "garut", "tasik"}
	var cities2 = cities[2:4]
	cities2[0] = "wado"

	fmt.Println("cities => ", cities)
	fmt.Println("cities2 => ", cities2)

	fmt.Println(strings.Repeat("#", 50))

	// cap function dan len function
	var cities3 = []string{"wado", "situraja", "sumedang", "darmaraja"}

	fmt.Println("cities3 cap:", cap(cities3)) //4
	fmt.Println("cities3 len:", len(cities3)) //4

	fmt.Println(strings.Repeat("#", 50))

	var cities4 = cities3[0:3]

	fmt.Println("cities4 cap:", cap(cities4)) //4
	fmt.Println("cities4 len:", len(cities4)) //3

	fmt.Println(strings.Repeat("#", 50))

	var cities5 = cities3[1:]

	fmt.Println("cities5 cap:", cap(cities5)) //3
	fmt.Println("cities5 len:", len(cities5)) //3

	fmt.Println(strings.Repeat("#", 50))

	// creating a new backing array / membuat slice dengan alokasi memory yang baru
	cars := []string{"Ford", "Honda", "Audhi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	// ketika slice cars melakukan replace element pada index 0 , tidak akan berdampak pada element slice newCars
	// dikarnakan variable slice cars dan newCars sudah berbeda alokasi memory nya
	cars[0] = "Nissan"
	fmt.Println("cars:", cars)
	fmt.Println("newCars:", newCars)
}
