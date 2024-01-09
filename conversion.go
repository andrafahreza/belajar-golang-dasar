package main

import "fmt"

func main() {
	var nilai32 int32 = 32123
	var nilai64 = int64(nilai32)

	fmt.Println(nilai64)

	var name = "Andra"
	var e = name[0]
	// uint to string
	var eString = string(e)

	fmt.Println(eString)
}
