package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func rot(x [][]int) [][]int {
	l := len(x)
	ret := make([][]int, l)
	for i := 0; i < l; i++ {
		ret[i] = make([]int, l)
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			ret[i][j] = x[l-1-j][i]
		}
	}
	return ret
}

func rotate(queue [][]int) [][]int {
	var rotated_queue [][]int

	return rotated_queue
}

func main() {
	var N int
	var sc = bufio.NewScanner(os.Stdin)

	fmt.Scan(&N)
	A, B := make([][]int, N), make([][]int, N)
	for i := 0; i < N; i++ {
		A[i] = make([]int, N)
		B[i] = make([]int, N)
	}

	for i := 0; i < 2*N; i++ {
		sc.Scan()
		aString := sc.Text()
		aStringList := strings.Split(aString, " ")
		for j, v := range aStringList {
			if i < N {
				A[i][j], _ = strconv.Atoi(v)
			} else {
				B[i-N][j], _ = strconv.Atoi(v)
			}
		}
	}

	for c := 0; c < 4; c++ {
		ok := true
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				if A[i][j] == 1 && B[i][j] != 1 {
					ok = false
				}
			}
		}
		if ok {
			fmt.Println("Yes")
			return
		}
		A = rot(A)
	}
	fmt.Println("No")
}
