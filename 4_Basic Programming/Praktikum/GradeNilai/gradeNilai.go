package main

import "fmt"

func main() {
	var nama string
	var nilai int

	fmt.Print("Masukkan nama siswa: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan nilai siswa: ")
	fmt.Scanln(&nilai)

	var deskripsi string

	switch {
	case nilai >= 80 && nilai <= 100:
		deskripsi = "A"
	case nilai >= 65 && nilai <= 79:
		deskripsi = "B"
	case nilai >= 50 && nilai <= 64:
		deskripsi = "C"
	case nilai >= 35 && nilai <= 49:
		deskripsi = "D"
	case nilai >= 0 && nilai <= 34:
		deskripsi = "E"
	default:
		deskripsi = "Nilai Invalid"
	}

	fmt.Printf("Nama siswa: %s\n", nama)
	fmt.Printf("Deskripsi nilai: %s\n", deskripsi)
}
