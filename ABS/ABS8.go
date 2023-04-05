package main

import "fmt"

func existCandidate(n int, total int) (int, int, int) {
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			if (10000*i+5000*j+1000*(n-i-j)) == total && n-i-j >= 0 {
				return i, j, n - i - j
			}
		}
	}

	return -1, -1, -1

}

func main() {
	var N, Y int
	fmt.Scan(&N, &Y)
	res1, res2, res3 := existCandidate(N, Y)
	fmt.Println(res1, res2, res3)
}
