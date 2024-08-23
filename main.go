package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Inisialisasi array jenis kopi
	jenisKopi := []string{"Kopi Expreso", "Kopi Latte", "Kopi Americano", "Kopi Cappuccino", "Kopi Mocca", "Kopi Campur Cuka"}

	// Set seed untuk generator nomor acak
	rand.Seed(time.Now().UnixNano())

	// Mainkan permainan sampai semua jenis kopi telah dicoba atau kopi yang tidak enak ditemukan
	for len(jenisKopi) > 0 {
		// Pilih jenis kopi secara acak dari array
		i := rand.Intn(len(jenisKopi))
		kopi := jenisKopi[i]

		// Cek jika kopi tidak enak
		if kopi == "Kopi Campur Cuka" {
			fmt.Printf("%s : gagal\n", kopi)
			break
		} else {
			fmt.Printf("%s : aman\n", kopi)
		}

		// Hapus jenis kopi dari array
		jenisKopi = append(jenisKopi[:i], jenisKopi[i+1:]...)
	}

	fmt.Println("Permainan selesai!")
}