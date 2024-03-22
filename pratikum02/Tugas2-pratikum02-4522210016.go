package main

import "fmt"

type Mahasiswa struct {
	Nama    string
	NPM     string
	Jurusan string
}

func main() {
	// Membuat map untuk menyimpan data mahasiswa
	dataMahasiswa := make(map[string]Mahasiswa)

	// Menambahkan data mahasiswa ke dalam map
	dataMahasiswa["4522210016"] = Mahasiswa{"Sulthan Zharfan", "4522210016", "Teknik Informatika"}
	dataMahasiswa["4522210002"] = Mahasiswa{"Dimas Stario", "4522210002", "Teknik Informatika"}
	dataMahasiswa["4522210002"] = Mahasiswa{"Udin Knalpot", "4522210899", "Parawisata"}
	dataMahasiswa["4522210002"] = Mahasiswa{"Asep Malik", "4522210455", "Psikolog"}
	dataMahasiswa["4522210002"] = Mahasiswa{"Malik Adam", "4522210232", "Manajemen"}

	// Menampilkan daftar nama mahasiswa dengan perulangan
	fmt.Println("Daftar Mahasiswa:")
	for _, mhs := range dataMahasiswa {
		fmt.Println(mhs.Nama)
	}

	// Mencari data mahasiswa berdasarkan NPM
	cariNPM := "4522210016"
	mhs, found := dataMahasiswa[cariNPM]
	if found {
		fmt.Printf("\nData Mahasiswa dengan NPM %s ditemukan:\n", cariNPM)
		fmt.Println("Nama:", mhs.Nama)
		fmt.Println("NPM:", mhs.NPM)
		fmt.Println("Jurusan:", mhs.Jurusan)
	} else {
		fmt.Printf("\nData Mahasiswa dengan NPM %s tidak ditemukan.\n", cariNPM)
	}
}
