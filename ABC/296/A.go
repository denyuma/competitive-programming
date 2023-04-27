package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)

	var flag bool = true

	for i := 1; i < n; i++ {
		if (s[i-1] == 'M' && s[i] == 'F') || (s[i-1] == 'F' && s[i] == 'M') {
			continue
		} else {
			flag = false
		}
	}

	if flag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
