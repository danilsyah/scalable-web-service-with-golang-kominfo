package helpers

import "fmt"

func TampilkanDataTeman(absen int) {
	found := false

	for _, teman := range DataTeman {
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
