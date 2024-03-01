package main

import (
	"fmt"
	"sync"
)

func main() {
	fruits := []string{"apple", "banana", "orange", "durian"}

	var wg sync.WaitGroup

	for index, fruit := range fruits {
		wg.Add(1)
		go printFruit(index, fruit, &wg)
	}
	wg.Wait()
	fmt.Println("proses main ended")
}

func printFruit(index int, fruit string, wg *sync.WaitGroup) {
	fmt.Printf("Index => %d, fuit => %s\n", index, fruit)
	wg.Done()
}
