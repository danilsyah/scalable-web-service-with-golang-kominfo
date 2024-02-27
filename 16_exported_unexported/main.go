package main

import "16_exported_unexported/helpers"

func main() {
	helpers.Greet()

	var person = helpers.Person{}

	person.Invokegreet()
}

// untuk menjalankan lebih dari 1 package main
// go run *.go
