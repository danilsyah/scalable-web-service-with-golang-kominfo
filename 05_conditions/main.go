package main

import "fmt"

func main() {

	var currentYear = 2024

	// if..else , if..else if..else
	// temporary variable age
	if age := currentYear - 1994; age < 17 {
		fmt.Println("kamu belum boleh membuat kartu SIM")
	} else {
		fmt.Println("kamu sudah boleh membuat kartu SIM")
	}

	// switch
	var score = 8

	switch score {
	case 8:
		fmt.Println("Perfect")
	case 7:
		fmt.Println("Awesome")
	case 6:
		fmt.Println("good")
	default:
		fmt.Println("not bad")
	}

	// switch with relational operators
	var nilai = 2

	switch {
	case nilai >= 8:
		fmt.Println("Perfect")
	case (nilai < 8) && (nilai > 3):
		fmt.Println("not bad")
	default:
		{
			fmt.Println("Study harder")
			fmt.Println("you need to learn more")
		}
	}

	// switch (fallthrough keyword)
	// melanjutkan pengecekan kepada case selanjutnya walaupun suatu case telah terpenuhi.
	score = 7

	switch {
	case score == 8:
		fmt.Println("perfect")
	case (score < 8) && (score > 3):
		fmt.Println("not bad euyy")
		fallthrough
	case score > 1:
		fmt.Println("It is Ok, but please study harder")
	default:
		{
			fmt.Println("study harder")
			fmt.Println("You don't have a good score ye")
		}
	}

	// nested conditions
	// pengkondisian bersarang
	score = 10

	if score > 7 {
		switch score {
		case 10:
			fmt.Println("perfect")
		default:
			fmt.Println("nice")
		}
	} else {
		if score == 5 {
			fmt.Println("not bad")
		} else if score == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if score == 0 {
				fmt.Println("try harder!")
			}
		}
	}

	var n int = 20

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizz Buzz", i)
		} else if i%3 == 0 {
			fmt.Println("Fizz", i)
		} else if i%5 == 0 {
			fmt.Println("Buzz", i)
		} else {
			fmt.Println(i)
		}
	}
}
