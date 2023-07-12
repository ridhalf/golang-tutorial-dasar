package main

import "fmt"

func main() {
	//memberikan alias untuk tipe data
	type noKtp string
	var ktp noKtp = "1234567890"
	fmt.Println(ktp)
}
