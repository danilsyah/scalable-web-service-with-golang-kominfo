package main

import "fmt"

func main() {
	// first way of looping
	for i := 0; i < 10; i++ {
		fmt.Printf("Angka %v \n", i)
	}

	// second way of looping
	var i = 0

	for i < 3 {
		fmt.Println("no ke - ", i)
		i++
	}

	// third way of looping
	var n = 0
	for {
		fmt.Println("Angka ke - ", n)
		n++
		if n == 3 {
			break
		}
	}

	// break and continue keywords
	for i := 1; i <= 10; i++ {
		/*
			memeriksa jika variable i memiliki nilai ganjil,
			maka proses loopingnya dipaksa berlanjut dan
			akan mengacuhkan syntax yang ada dibawah keyword continue
		*/
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("angka genap => ", i)
	}

	// cetak angka deret (nested looping)
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Print("\n")
	}

	// nested  looping using continue
	for i := 0; i <= 3; i++ {
		if i%2 == 0 {
			fmt.Println("Perulangan Genap ke - ", i)
			for j := i; j < 10; j++ {
				if j%2 == 1 {
					continue
				}
				fmt.Print(j, " ")
			}
			fmt.Print("\n")
		} else {
			fmt.Println("Perulangan Ganjil ke - ", i)
			for j := i; j < 10; j++ {
				if j%2 == 0 {
					continue
				}
				fmt.Print(j, " ")
			}
			fmt.Print("\n")
		}
	}

	// looping using label
outerLoop:
	for i := 0; i < 3; i++ {
		fmt.Println("Perulangan ke ", i+1)
		for j := 0; j < 3; j++ {
			if i == 2 {
				break outerLoop
			}
			fmt.Print(j, " ")
		}
		fmt.Print("\n")
	}
}
