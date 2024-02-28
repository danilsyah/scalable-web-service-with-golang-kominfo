package main

import (
	"assigment_1/helpers"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Gunakan: go run biodata.go [nomor_absen]")
		return
	}

	absen := os.Args[1]
	var absenInt int
	_, err := fmt.Sscanf(absen, "%d", &absenInt)
	if err != nil {
		fmt.Println("Nomor absen harus berupa bilangan bulat.")
		return
	}
	helpers.TampilkanDataTeman(absenInt)
}
