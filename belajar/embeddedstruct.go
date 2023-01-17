package main

import "fmt"

type person struct { // mendeklarasikan struct
	name string
	age  int
}

type student struct { // mendeklarasikan struct
	grade  int
	person // mengembed struct yang sudah dibuat
}

func main() {
	var s1 = student{}
	s1.name = "avatar"
	s1.age = 22
	s1.grade = 3

	fmt.Println("name  :", s1.name)
	fmt.Println("age   :", s1.age)
	fmt.Println("age   :", s1.person.age)
	fmt.Println("grade :", s1.grade)
}
