package main

import "fmt"

func main() {
	var nilai1 int32 = 3000
	var nilai2 int64 = int64(nilai1)
	fmt.Println(nilai2)

	var name = "Ridhal"
	var e byte = name[0]
	var eString string = string(e)

	fmt.Println(e)
	fmt.Println(eString)
}
