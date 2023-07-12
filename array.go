package main

import "fmt"

func main() {
	var names [2]string
	names[0] = "Ridhal"
	names[1] = "Fajri"

	fmt.Println(names[0])
	fmt.Println(names[1])

	var values = [3]int{
		1, 2, 3,
	}
	fmt.Println(values)
}
