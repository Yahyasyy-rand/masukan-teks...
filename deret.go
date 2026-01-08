package main

import (
	"fmt"
)

// Soal 5: Relasi Rekurens Iteratif
func soal5Deret(C1, C2 int, N int) {
	fmt.Println("=== Soal 5: Relasi Rekurens Iteratif ===")
	fmt.Printf("Parameter Paket 17: C1 = %d, C2 = %d, N = %d\n", C1, C2, N)

	// Kondisi awal
	a0 := 1
	a1 := 1

	// Jika N adalah 0 atau 1
	if N == 0 {
		fmt.Printf("HASIL AKHIR Suku ke=%d: %d\n\n", N, a0)
		return
	}
	if N == 1 {
		fmt.Printf("HASIL AKHIR Suku ke=%d: %d\n\n", N, a1)
		return
	}

	// Hitung iteratif
	fmt.Print("Process Perhitungan:\n")
	fmt.Printf("Suku 0: %d | ", a0)
	fmt.Printf("Suku 1: %d | ", a1)

	prev2 := a0 // a_{n-2}
	prev1 := a1 // a_{n-1}
	current := 0

	for i := 2; i <= N; i++ {
		current = C1*prev1 + C2*prev2

		if i <= 10 || i == N { // Tampilkan beberapa suku saja
			fmt.Printf("Suku %d: %d | ", i, current)
			if i%5 == 0 {
				fmt.Println()
			}
		}

		// Geser untuk iterasi berikutnya
		prev2 = prev1
		prev1 = current
	}

	fmt.Printf("\nHASIL AKHIR Suku ke=%d: %d\n\n", N, current)
}

// Soal 6: Analisis Kedekatan Deret Geometri
func soal6Deret(a float64, r float64, N int) {
	fmt.Println("=== Soal 6: Analisis Kedekatan Deret Geometri ===")
	fmt.Printf("Parameter Paket 17: a = %.1f, r = %.1f, N = %d\n", a, r, N)

	// Hitung jumlah N suku berhingga (S_N)
	// S_N = a * (1 - r^N) / (1 - r) jika r != 1
	var Sn float64

	if r == 1 {
		Sn = a * float64(N)
	} else {
		rToN := 1.0
		for i := 0; i < N; i++ {
			rToN *= r
		}
		Sn = a * (1 - rToN) / (1 - r)
	}

	// Hitung jumlah tak hingga (S)
	// Hanya valid jika |r| < 1
	var Sinf float64
	validInfinite := false

	if r > -1 && r < 1 {
		Sinf = a / (1 - r)
		validInfinite = true
	}

	// Hitung persentase kedekatan
	var percentage float64
	if validInfinite {
		percentage = (Sn / Sinf) * 100
	}

	fmt.Printf("Sum Berhingga S(%d): %.2f\n", N, Sn)

	if validInfinite {
		fmt.Printf("Sum Tak Hingga S: %.2f\n", Sinf)
		fmt.Printf("Persentase Kedekatan: %.2f%%\n\n", percentage)
	} else {
		fmt.Println("Sum Tak Hingga: Tidak terdefinisi (|r| >= 1)")
		fmt.Printf("Persentase Kedekatan: Tidak dapat dihitung\n\n")
	}
}
