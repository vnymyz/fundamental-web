package main

import "fmt"

type mahasiswa struct { // mendeklarasikan struct
	nama  string
	nilai int
}

func main() {
	var s1 mahasiswa // s1 adalah nama variabel dan mahasiswa adalah struct nya jadi diikuti gtu
	s1.nama = "vanya"
	s1.nilai = 8

	fmt.Println("nama anda :", s1.nama)   // vanya
	fmt.Println("nilai anda :", s1.nilai) // 8
}

// contoh codingan untuk nomor 2
