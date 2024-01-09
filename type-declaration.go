package main

import "fmt"

func main() {
	// ini alias
	type noKtp string

	var ktp noKtp = "123123123"
	fmt.Println(ktp)
	fmt.Println(noKtp("33333333"))
}
