package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Fungsi untuk menghapus duplikat dari slice
func removeDuplicates(arr []int) []int {
	seen := make(map[int]bool)
	result := []int{}

	for _, val := range arr {
		if !seen[val] {
			seen[val] = true
			result = append(result, val)
		}
	}
	return result
}

// Fungsi untuk mengecek apakah elemen ada dalam slice
func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// Fungsi operasi himpunan
func symmetricDifference(A, B []int) []int {
	result := []int{}

	// Elemen di A tapi tidak di B
	for _, a := range A {
		if !contains(B, a) {
			result = append(result, a)
		}
	}

	// Elemen di B tapi tidak di A
	for _, b := range B {
		if !contains(A, b) {
			result = append(result, b)
		}
	}

	return removeDuplicates(result)
}

func difference(A, B []int) []int {
	result := []int{}

	for _, a := range A {
		if !contains(B, a) {
			result = append(result, a)
		}
	}

	return removeDuplicates(result)
}

func intersection(A, B []int) []int {
	result := []int{}

	for _, a := range A {
		if contains(B, a) {
			result = append(result, a)
		}
	}

	return removeDuplicates(result)
}

func union(A, B []int) []int {
	result := append([]int{}, A...)
	result = append(result, B...)
	return removeDuplicates(result)
}

// Soal 1: Operasi Himpunan Kompleks
func soal1Himpunan(N int) {
	fmt.Println("=== Soal 1: Operasi Himpunan Kompleks ===")
	fmt.Printf("Parameter Paket 17: N = %d\n", N)

	// Generate himpunan A, B, C secara acak
	rand.Seed(time.Now().UnixNano())

	// Ukuran himpunan acak antara 3-5
	sizeA := rand.Intn(3) + 3
	sizeB := rand.Intn(3) + 3
	sizeC := rand.Intn(2) + 2

	A := make([]int, sizeA)
	B := make([]int, sizeB)
	C := make([]int, sizeC)

	// Generate elemen unik
	for i := 0; i < sizeA; i++ {
		A[i] = rand.Intn(N) + 1
	}
	for i := 0; i < sizeB; i++ {
		B[i] = rand.Intn(N) + 1
	}
	for i := 0; i < sizeC; i++ {
		C[i] = rand.Intn(N) + 1
	}

	// Hapus duplikat
	A = removeDuplicates(A)
	B = removeDuplicates(B)
	C = removeDuplicates(C)

	fmt.Printf("A: %v | B: %v | C: %v\n", A, B, C)

	// Operasi: R
	// 1. A B perbedaan simetrik (symmetric difference)
	A_sym_B := symmetricDifference(A, B)
	fmt.Printf("1. A B = %v\n", A_sym_B)

	// 2. (A B) \ C
	diff := difference(A_sym_B, C)
	fmt.Printf("2. (A B) \\ C = %v\n", diff)

	// 3. A irisan C
	A_int_C := intersection(A, C)
	fmt.Printf("3. A irisan C = %v\n", A_int_C)

	// 4. Gabungan: R = {2} ∪ {3}
	R := union(diff, A_int_C)
	fmt.Printf("4. R = %v ∪ %v = %v\n", diff, A_int_C, R)

	// 5. Filter: Bilangan Genap dan < (N/4)
	filterLimit := N / 4
	filtered := []int{}
	for _, val := range R {
		if val%2 == 0 && val < filterLimit {
			filtered = append(filtered, val)
		}
	}

	fmt.Printf("5. Filter (Genap < %d): %v\n", filterLimit, filtered)
	fmt.Printf("Total Elemen: %d\n\n", len(filtered))
}

// Soal 2: Analisis Pasangan Subset
func soal2Himpunan(M, K int) {
	fmt.Println("=== Soal 2: Analisis Pasangan Subset ===")
	fmt.Printf("Parameter Paket 17: M = %d, K = %d\n", M, K)

	// Generate himpunan S dengan M elemen unik
	rand.Seed(time.Now().UnixNano())
	S := make([]int, M)

	// Generate elemen unik antara 1-50
	used := make(map[int]bool)
	for i := 0; i < M; i++ {
		for {
			val := rand.Intn(50) + 1
			if !used[val] {
				S[i] = val
				used[val] = true
				break
			}
		}
	}

	fmt.Printf("Set S: %v | Target K: %d\n", S, K)

	// Cari semua pasangan unik {x, y} dimana x != y
	pairs := [][2]int{}

	for i := 0; i < len(S); i++ {
		for j := i + 1; j < len(S); j++ {
			pairs = append(pairs, [2]int{S[i], S[j]})
		}
	}

	// Filter pasangan dengan jumlah > K
	filteredPairs := [][2]int{}
	for _, pair := range pairs {
		if pair[0]+pair[1] > K {
			filteredPairs = append(filteredPairs, pair)
		}
	}

	fmt.Println("Subset 2-Elemen (Sum > K):")
	for i, pair := range filteredPairs {
		fmt.Printf("%d. {%d, %d} (Sum=%d)\n", i+1, pair[0], pair[1], pair[0]+pair[1])
	}
	fmt.Printf("Total: %d Pasangan\n\n", len(filteredPairs))
}
