package main

import "fmt"

func main() {
	var a, b, c int
	var s string

	fmt.Scan(&a)
	fmt.Scan(&b, &c)
	fmt.Scan(&s)
	sum := a + b + c

	fmt.Printf("%d %s\n", sum, s)
}
