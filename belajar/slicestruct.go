package main

import "fmt"

type person struct { // mendeklarasikan struct
	name string
	age  int
}

func main() { // membuat function main
	var allStudents = []person{ // person membuat slice
		{name: "Wick", age: 23}, // properti yang diisi sesuai struct
		{name: "Ethan", age: 23},
		{name: "Bourne", age: 22},
	}

	for _, student := range allStudents { // perulangan sesuai dengan range all student
		fmt.Println(student.name, "age is", student.age) // cetak sesuai struct
	}
}

// contoh codingan untuk nomor 5
