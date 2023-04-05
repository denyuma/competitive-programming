package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)
	var t = make([]int, N+1)
	var x = make([]int, N+1)
	var y = make([]int, N+1)
	t[0] = 0
	x[0] = 0
	y[0] = 0
	for i := 0; i < N; i++ {
		fmt.Scan(&t[i+1], &x[i+1], &y[i+1])
	}

	flag := true

	for i := 0; i < N; i++ {
		dt := t[i+1] - t[i]
		dist := math.Abs(float64(x[i+1]-x[i])) + math.Abs(float64(y[i+1]-y[i]))
		if int(dist) > dt {
			flag = false
		}
		if int(dist)%2 != dt%2 {
			flag = false
		}
	}

	if flag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
