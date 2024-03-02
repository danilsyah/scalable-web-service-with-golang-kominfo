package main

import (
	"errors"
	"fmt"
)

func main() {
	/*
		Panic digunakan untuk menampilkan stack traceerror sekaligus menghentikan flow goroutine(karena main() juga merupakan goroutine, makabehaviour yang sama juga berlaku).
	*/
	defer catchErr()

	var password string

	fmt.Printf("Input Password : ")
	fmt.Scanln(&password)

	if valid, err := validPassword1(password); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(valid)
	}

}

func catchErr() {
	if r := recover(); r != nil {
		fmt.Println("Error occured:", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func validPassword1(password string) (string, error) {
	pl := len(password)

	if pl < 5 {
		return "", errors.New("password has to have more than 4 characters")
	}

	return "Valid password", nil
}
