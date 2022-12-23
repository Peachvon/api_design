package main

import "fmt"

type Math func(int, int) int

func cal(sn Math) int {
	return sn(5, 4)
}

func sum(a int, b int) int {
	return a + b
}

func main() {
	fn := sum
	r1 := fn(1, 2)
	fmt.Printf("fn(1,2): %v ln", r1)

	r2 := cal(fn)
	fmt.Printf("cal(fn): %v", r2)

	r3 := cal(sum)
	fmt.Printf("cal(sum): %v", r3)

}
