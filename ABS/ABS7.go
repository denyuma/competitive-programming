package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)
	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	text := sc.Text()
	strList := strings.Split(text, " ")
	intList := make([]int, len(strList))
	for i, v := range strList {
		intList[i], _ = strconv.Atoi(v)
	}

	sort.Ints(intList)
	sort.Sort(sort.Reverse(sort.IntSlice(intList)))

	alicePoint := 0
	bobPoint := 0

	for i, v := range intList {
		if i%2 == 0 {
			alicePoint += v
		} else {
			bobPoint += v
		}
	}

	result := alicePoint - bobPoint
	fmt.Println(result)
}
