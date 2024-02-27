package main

import "fmt"

/*
	terdapat suatu function yang akan dieksekusi terlebih dahulu sebelum function main. Function tersebut bernama init
*/
func init() {
	fmt.Println("Hallo!!! Ini berasal dari function init")
}
