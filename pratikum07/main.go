package main

import (
	"errors"
	"fmt"
)

// Definisi error
var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

// Fungsi untuk mendapatkan ID
func GetByid(id string) error {
	if id == "" { // Menggunakan id yang kosong sebagai error
		return ValidationError
	}

	if id != "zharfan" { // Jika ID tidak sama dengan "zharfan", kembalikan NotFoundError
		return NotFoundError
	}

	// sukses
	return nil
}

func main() {
	// Memeriksa apakah ada kesalahan saat mendapatkan ID
	err := GetByid("zharfan")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("not found error")
		} else {
			fmt.Println("unknown error")
		}
	}
}
