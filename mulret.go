package main

import "fmt"

func vals() (int, int) {
	return 1, 2
}

func hello() (string, string) {
	return "Hello!", "World"
}

func main() {
	a, b := vals()
	fmt.Println(a, b)
	fmt.Println(b)
	fmt.Println(hello())
}
