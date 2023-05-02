package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i == 4 {
			continue
		}
		fmt.Println(2 * i)

	}
}
