package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Menangani permintaan HTTP di rute "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Memeriksa metode HTTP yang digunakan dalam permintaan
		switch r.Method {
		case "POST":
			w.Write([]byte("post")) // Menulis respons untuk metode POST
		case "GET":
			w.Write([]byte("get")) // Menulis respons untuk metode GET
		default:
			http.Error(w, "Metode tidak didukung", http.StatusBadRequest) // Penanganan metode lain
		}
	})

	// Menjalankan server pada port 9000
	fmt.Println("Server started at localhost:9000")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		fmt.Println("Server error:", err) // Penanganan kesalahan saat menjalankan server
	}
}
