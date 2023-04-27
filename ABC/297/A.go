package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N, D int
	var flag int = -1
	fmt.Scan(&N, &D)
	T := make([]int, N)
	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	tString := sc.Text()
	tStringList := strings.Split(tString, " ")
	for i, v := range tStringList {
		T[i], _ = strconv.Atoi(v)
	}

	for i := 0; i < N-1; i++ {
		diff := T[i+1] - T[i]
		if diff <= D {
			flag = T[i+1]
			break
		}
	}

	fmt.Println(flag)
}
