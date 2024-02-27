package main

import (
	"fmt"
	"os"
)

// Teman adalah struct untuk menyimpan informasi teman di kelas
type Teman struct {
	Absen     int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// dataTeman adalah slice untuk menyimpan data teman
var dataTeman = []Teman{
	{1, "John", "Jakarta", "Software Engineer", "Ingin belajar pemrograman Go"},
	{2, "Doe", "Bandung", "Data Scientist", "Tertarik dengan kecepatan dan kesederhanaan Go"},
	// Tambahkan data teman lainnya di sini
}

// tampilkanDataTeman adalah function untuk menampilkan informasi teman berdasarkan nomor absen
func tampilkanDataTeman(absen int) {
	found := false
	for _, teman := range dataTeman {
		if teman.Absen == absen {
			fmt.Printf("Nama: %s\n", teman.Nama)
			fmt.Printf("Alamat: %s\n", teman.Alamat)
			fmt.Printf("Pekerjaan: %s\n", teman.Pekerjaan)
			fmt.Printf("Alasan memilih kelas Golang: %s\n", teman.Alasan)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Teman dengan nomor absen tersebut tidak ditemukan.")
	}
}

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

	tampilkanDataTeman(absenInt)
}
