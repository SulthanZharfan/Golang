package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Membuat scanner untuk input dari konsol
	scanner := bufio.NewScanner(os.Stdin)

	// Input nama
	fmt.Print("Masukkan nama Anda: ")
	scanner.Scan()
	nama := scanner.Text()

	// Input usia
	fmt.Print("Masukkan usia Anda: ")
	scanner.Scan()
	var usia int
	fmt.Sscanf(scanner.Text(), "%d", &usia)

	// Menentukan kategori usia
	var kategoriUsia string
	if usia < 18 {
		kategoriUsia = "anak-anak"
	} else {
		kategoriUsia = "dewasa"
	}

	// Menampilkan pesan selamat datang dan kategori usia
	fmt.Printf("Selamat datang, %s! Anda termasuk kategori %s.\n", nama, kategoriUsia)
}
