// nomor 2
// output tidak harus sesuai, yang penting menerapkan semua dibawah ini
package main

import "fmt"

type film struct { // mendeklarasikan struct
	judul string
	genre string
	tahun string
	nilai int
}

func main() {
	var p1 film
	p1.judul = "avatar"
	p1.genre = "fantasy"
	p1.tahun = "2023"
	p1.nilai = 9

	fmt.Println("Judul Film   		:", p1.judul)
	fmt.Println("Genre Film   		:", p1.genre)
	fmt.Println("Tahun film release 	:", p1.tahun)
	fmt.Println("Nilai film 		:", p1.nilai)
}

// contoh codingan untuk nomor 2
