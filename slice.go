package main

import "fmt"

func main() {
	var month = [...]string{
		"januari",
		"februari",
		"maret",
		"april", "mei", "juni", "juli", "agustrus", "september", "oktober", "november",
	}
	var slice1 = month[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
}
