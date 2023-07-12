package main

import "fmt"

func main() {
	var name string
	name = "Ridhal Fajri"
	fmt.Println(name)

	name = "Ridhal"
	fmt.Println(name)

	age := 22 // mengganti deklarasi variabel menggunakan var
	fmt.Println(age)

	//dikelompokkan dalam deklarasi variabel
	var (
		firstName = "Ridhal"
		lastName  = "Fajri"
	)
	fmt.Println(firstName, lastName)
}
