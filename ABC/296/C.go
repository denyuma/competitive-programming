package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const MaxBuf = 2001000

var buf []byte = make([]byte, MaxBuf)
var sc = bufio.NewScanner(os.Stdin)

func readInt() int {
	sc.Scan()
	r, _ := strconv.Atoi(sc.Text())
	return r
}

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(buf, MaxBuf)
}

func main() {
	n := readInt()
	x := readInt()

	a := make(map[int]struct{}, n)
	for i := 0; i < n; i++ {
		v := readInt()
		a[v] = struct{}{}
	}

	for k := range a {
		if _, ok := a[k+x]; ok {
			fmt.Println("Yes")
			os.Exit(0)
		}
	}

	fmt.Println("No")
	os.Exit(0)
}
