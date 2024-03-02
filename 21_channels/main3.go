package main

import (
	"fmt"
	"time"
)

func main() {
	// Buffered channel yang dimana kita dapat menentukkan kapasitas dari buffer nya, dan selama jumlah data yang dikirim melalui buffered channel tidak melebihi kapasitasnya, maka proses pengiriman data akan bersifat asynchronous.

	c1 := make(chan int, 3)

	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Printf("func goroutine #%d starts sending data into the channel\n", i)
			c <- i
			fmt.Printf("func goroutine #%d after sending data into the channel\n", i)
		}
		close(c)
	}(c1)

	fmt.Println("main goroutine sleeps 2 seconds")
	time.Sleep(time.Second * 2)

	for v := range c1 { // v = <- c1
		fmt.Println("main goroutine received value from channel:", v)
	}

}
