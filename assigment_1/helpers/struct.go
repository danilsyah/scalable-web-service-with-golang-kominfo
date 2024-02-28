package helpers

// Teman adalah struct untuk menyimpan informasi data teman
type Teman struct {
	Absen     int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// DataTeman adalah slice untuk menyimpan data teman
var DataTeman = []Teman{
	{1, "Danil", "Bandung", "Software Engineer", "Ingin belajar pemrograman Go"},
	{2, "Haykal", "Sumedang", "Backend Engineer", "Tertarik dengan kecepatan dan kesederhanaan Go"},
	{3, "Nufika", "Papua", "Data Analyst", "Go tanpa OOP"},
	// Tambahkan data teman lainnya di sini
}
