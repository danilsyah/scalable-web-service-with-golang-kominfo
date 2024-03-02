package main

import (
	"fmt"
	"time"
)

func main() {

	// UnBuffered Channel = dimana proses penerimaan dan pengiriman data bersifat synchronous.
	// syntax apapun yang berada dibawah proses pengiriman data melalui unbuffered channel akan tertahan hingga datanya diterima oleh Goroutine lainnya
	c1 := make(chan int)

	go func(c chan int) {
		fmt.Println("func goroutine starts sending data into the channel")
		c <- 10
		fmt.Println("func goroutine after sending data into the channel")
	}(c1)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	fmt.Println("main goroutine start receiving data")
	d := <-c1
	fmt.Println("main goroutine received data:", d)

	close(c1)
	time.Sleep(time.Second)
}
