package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
	Arti dari concurrency adalah mengeksekusi sebuah proses secara independen atau berhadapan dengan lebih dari satu tugas dalam waktu yang sama. Perlu diingat disini bahwa concurrency  berbeda dengan parallelism, karena parallelism memiliki arti mengerjakan tugas yang banyak secara bersamaan dari awal hingga akhir. Sedangkan pada concurrency, kita tidak akan tahu tentang siapa yang akan menyelesaikan tugasnya terlebih dahulu.

*/
func main() {
	/*
		Goroutine bersifat asynchronous sehingga proses nya tidak saling tunggu dengan Goroutine lainnya
	*/
	fmt.Println("main execution started")

	go firstProcess(8)

	secondProcess(8)

	fmt.Println("No. of Goroutines:", runtime.NumGoroutine())

	time.Sleep(time.Second * 2)

	fmt.Println("Main execution ended")

}

func firstProcess(index int) {
	fmt.Println("First process func started")
	for i := 1; i <= index; i++ {
		fmt.Println("i=", i)
	}
	fmt.Println("First process func ended")
}

func secondProcess(index int) {
	fmt.Println("Second process func started")
	for j := 1; j <= index; j++ {
		fmt.Println("j=", j)
	}
	fmt.Println("Second process func ended")
}
