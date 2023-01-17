// nomor 5
// output tidak harus sesuai yang penting menerapkan semua yang dibawah
package main

import "fmt"

type makanan struct { // mendeklarasikan struct
	nama   string
	negara string
	jumlah int
}

func main() { // membuat function main
	var semuaMakanan = []makanan{ // person membuat slice
		{nama: "Pizza", negara: "Italy", jumlah: 80}, // properti yang diisi sesuai struct
		{nama: "Ramen", negara: "Jepang", jumlah: 100},
		{nama: "Rendang", negara: "Indonesia", jumlah: 200},
	}

	for _, laper := range semuaMakanan { // perulangan sesuai dengan range all student
		fmt.Println(laper.nama, "berasal dari", laper.negara, "yang tersedia di restoran berjumlah", laper.jumlah) // cetak sesuai struct
	}
}

// contoh codingan untuk nomor 5
