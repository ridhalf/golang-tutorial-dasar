package main

import "fmt"

func main() {
	var name = "Tom"
	if name == "Ridhal" {
		fmt.Println("Halo", name)
	} else if name == "Tom" {
		fmt.Println("Hi,", name)
	} else {
		fmt.Println("Hi, Boleh kenalan?")
	}

	//short statement pada if
	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
