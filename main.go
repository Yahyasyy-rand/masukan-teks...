package main

import "fmt"

func main() {
	fmt.Println("UTS Pemrograman Golang Matematika Diskrit - Paket 17")
	fmt.Println("STMIK Tazkia - Prodi Teknik Informatika - Semester Ganjil 2025/2026")
	fmt.Println("Oleh: Muhammad Yahya Ayyasy dan Rizwan Defriansyah")
	fmt.Println("NIM : 251552010054 dan 251552010002")
	fmt.Println("===========================================\n")

	// Parameter Paket 17 dari tabel
	// Pkt 17: 55, 7, 10, 3, 3, 4, -2, 9, 3, 0.6, 9
	// Kolom: S1(N), S2(M,K), S3(N), S4(N), S5(C1,C2,N), S6(a,r,N)

	// Materi I: Himpunan
	soal1Himpunan(55)    // N = 55
	soal2Himpunan(7, 10) // M = 7, K = 10

	// Materi II: Matriks
	soal3Matrix(3) // N = 3 (ordo matriks untuk soal 3)
	soal4Matrix(3) // N = 3 (ordo matriks untuk soal 4)

	// Materi III: Fungsi dan Deret
	soal5Deret(4, -2, 9)  // C1 = 4, C2 = -2, N = 9
	soal6Deret(3, 0.6, 9) // a = 3, r = 0.6, N = 9

	fmt.Println("===========================================")
	fmt.Println("Program selesai.")
}
