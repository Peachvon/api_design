package main

import "fmt"

type Decorator func(s string) error

func Use(next Decorator) Decorator {
	return func(c string) error {
		c = c + " peach"
		return next(c)
	}
}

func home(s string) error {
	fmt.Println("home", s)
	return nil
}

func main() {
	wrapped := Use(home)

	w := wrapped("world")
	h := wrapped("world")
	fmt.Println("wrapped", w)
	fmt.Println("wrapped", h)
}
