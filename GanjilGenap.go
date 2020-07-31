package main

import "fmt"

func main() {
	angka := 1002

	if angka%2 == 0 {
		fmt.Printf("%d adalah bilangan genap", angka)
	} else {
		fmt.Printf("%d adalah bilangan ganjil", angka)
	}
	// nothing wrong, hanya memberi cara lain yang mungkin lebih simple
}
