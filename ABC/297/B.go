package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	fmt.Scan(&S)
	x1 := strings.Index(S, "B")
	x2 := strings.LastIndex(S, "B")
	y1 := strings.Index(S, "R")
	y2 := strings.LastIndex(S, "R")
	z := strings.Index(S, "K")

	var flag bool = false
	if (x1%2 != x2%2) && (y1 < z && z < y2) {
		flag = true
	}

	if flag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
