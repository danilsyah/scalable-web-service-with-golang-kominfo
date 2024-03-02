package main

import (
	"fmt"
	"os"
)

/*
Funcion exit yang yang berasal dari package os berguna untuk menghentikan suatu program secara paksa.
*/
func main() {
	defer fmt.Println("Invoke with defer")
	fmt.Println("Before Exiting")
	os.Exit(1)
}
