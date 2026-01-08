package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Soal 3: Perkalian dan Trace Matriks
func soal3Matrix(N int) {
	fmt.Println("=== Soal 3: Perkalian dan Trace Matriks ===")
	fmt.Printf("Parameter Paket 17: N = %d (Ordo Matriks)\n", N)

	// Generate matriks A dan B
	rand.Seed(time.Now().UnixNano())

	// Buat matriks A
	A := make([][]int, N)
	for i := range A {
		A[i] = make([]int, N)
		for j := range A[i] {
			A[i][j] = rand.Intn(10) + 1 // Angka 1-10
		}
	}

	// Buat matriks B
	B := make([][]int, N)
	for i := range B {
		B[i] = make([]int, N)
		for j := range B[i] {
			B[i][j] = rand.Intn(10) + 1 // Angka 1-10
		}
	}

	fmt.Printf("Matrix A (%dx%d):\n", N, N)
	printMatrix(A)

	fmt.Printf("Matrix B (%dx%d):\n", N, N)
	printMatrix(B)

	// Perkalian matriks R = A × B
	R := make([][]int, N)
	for i := range R {
		R[i] = make([]int, N)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			sum := 0
			for k := 0; k < N; k++ {
				sum += A[i][k] * B[k][j]
			}
			R[i][j] = sum
		}
	}

	fmt.Printf("Hasil Matriks R (%dx%d):\n", N, N)
	printMatrix(R)

	// Hitung trace (jumlah diagonal utama)
	trace := 0
	for i := 0; i < N; i++ {
		trace += R[i][i]
	}

	fmt.Printf("Trace: %d\n\n", trace)
}

// Soal 4: Transformasi Baris
func soal4Matrix(N int) {
	fmt.Println("=== Soal 4: Transformasi Baris ===")
	fmt.Printf("Parameter Paket 17: N = %d (Ordo Matriks)\n", N)

	// Generate matriks M
	rand.Seed(time.Now().UnixNano())

	M := make([][]int, N)
	for i := range M {
		M[i] = make([]int, N)
		for j := range M[i] {
			M[i][j] = rand.Intn(20) + 1 // Angka 1-20
		}
	}

	fmt.Printf("Matrix M (%dx%d) Generated:\n", N, N)
	printMatrix(M)

	// Tukar baris 0 dengan baris N-1
	if N > 1 {
		M[0], M[N-1] = M[N-1], M[0]
		fmt.Println("Menukar Baris 0 dan", N-1, "...")
		fmt.Printf("Matrix M Terkini:\n")
		printMatrix(M)
	}

	// Cari nilai maksimum dan posisinya
	maxVal := M[0][0]
	maxRow, maxCol := 0, 0

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if M[i][j] > maxVal {
				maxVal = M[i][j]
				maxRow = i
				maxCol = j
			}
		}
	}

	fmt.Printf("Nilai Maksimum: %d ditemukan di Posisi (%d,%d)\n\n", maxVal, maxRow, maxCol)
}

// Fungsi bantu untuk mencetak matriks
func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Print("[")
		for j, val := range row {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Printf("%2d", val)
		}
		fmt.Println("]")
	}
}
