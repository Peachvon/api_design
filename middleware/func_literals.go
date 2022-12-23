package main

import "fmt"

func main() {
	r := func(a, b int) bool {
		return a < b
	}(4, 3)

	fmt.Println(r)
}
