package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Ridhal",
		"address": "Bogor",
	}
	person["title"] = "Programmer"
	fmt.Println(person)

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Ridhal"
	book["ups"] = "salah"
	fmt.Println(book)
	fmt.Println(len(book))
	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))

}
