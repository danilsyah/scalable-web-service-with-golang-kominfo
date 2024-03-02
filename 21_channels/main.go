/*
Channel merupakan sebuah mekanisme untuk Goroutine saling berkomunikasi dengan Goroutine lainnya. Maksud berkomunikasi disini adalah proses pengiriman dan pertukaran data antara Goroutine satu dengan Goroutine lainnya.
*/
package main

import "fmt"

func main() {
	c := make(chan string)

	students := []string{"danil", "haykal", "fika"}

	for _, v := range students {
		go introduce(v, c)
	}

	for i := 1; i <= 3; i++ {
		print(c)
	}

	close(c)

}

func print(c <-chan string) {
	fmt.Println(<-c)
}

func introduce(student string, c chan<- string) {
	result := fmt.Sprintf("hai, my name is %s", student)
	c <- result
}
