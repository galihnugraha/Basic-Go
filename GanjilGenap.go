package main

import (
	"fmt"
)

func main() {
	var angka = 1002

	if angka%2 == 0 {
		fmt.Println("angka tersebut merupakan bilangan genap")
	}

	if angka%2 == 1 {
		fmt.Println("angka tersebut merupakan bilangan ganjil")
	}
}
