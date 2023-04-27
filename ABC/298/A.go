package main

import "fmt"

func main() {
	var N int
	var S string

	fmt.Scan(&N)
	fmt.Scan(&S)

	var flag_one, flag_two bool = false, false

	for _, v := range S {
		if v == 'o' {
			flag_one = true
		} else if v == 'x' {
			flag_two = true
		}
	}

	if flag_one && !flag_two {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
