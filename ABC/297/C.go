package main

import (
	"fmt"
)

func main() {
	var H, W int
	fmt.Scan(&H, &W)

	S := make([][]byte, H)
	for i := range S {
		S[i] = make([]byte, W)
		for j := 0; j < W; j++ {
			fmt.Scanf("%c", &S[i][j])
		}

		fmt.Scanf("%*c")

		for j := 0; j < W-1; {
			if S[i][j] == 'T' && S[i][j+1] == 'T' {
				S[i][j] = 'P'
				S[i][j+1] = 'C'
				j += 2
			} else {
				j += 1
			}
		}

		for j := range S[i] {
			fmt.Printf("%c", S[i][j])
		}
		fmt.Println("")

	}

}
