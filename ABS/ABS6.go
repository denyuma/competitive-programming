package main

import "fmt"

func calcSumOfDigits(num int) (int, int) {
	copyOfNum := num
	sum := 0

	for {
		sum += (copyOfNum % 10)
		copyOfNum /= 10
		if copyOfNum == 0 {
			break
		}
	}

	return num, sum
}

func main() {
	var N, A, B int
	var result int = 0

	fmt.Scan(&N, &A, &B)

	for i := 0; i <= N; i++ {
		num, sum := calcSumOfDigits(i)
		if A <= sum && sum <= B {

			result += num
		}
	}

	fmt.Println(result)
}
