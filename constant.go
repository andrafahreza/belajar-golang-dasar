package main

import "fmt"

// ================ Catatan ==============

// Hampir sama dengan variable, tetapi value tidak bisa diubah

// ================	Catatan	===============

func main() {
	const firstName string = "Andra"
	const lastName = "Fahreza"

	fmt.Println(firstName)
	fmt.Println(lastName)

	// Multiple
	const (
		first  string = "pertama"
		second int    = 2
	)
	fmt.Println(first)
	fmt.Println(second)
}
