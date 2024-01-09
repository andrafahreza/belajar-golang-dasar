package main

import "fmt"

func main() {
	// Dengan tipe data
	var name string

	name = "Andra Fahreza"
	fmt.Println(name)

	name = "Ubah Nama"
	fmt.Println(name)

	// Tidak dengan tipe data
	var name2 = "Andra Fahreza"
	fmt.Println(name2)

	// Dengan Var
	awal := "Tes ubah"
	fmt.Println(awal)

	// Multiple variable
	var (
		firstName = "Mhd Andra"
		lastName  = "Fahreza"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)

}
