package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertStringListToIntList(data []string) []int {
	intList := make([]int, len(data))
	for i, v := range data {
		intList[i], _ = strconv.Atoi(v)
	}
	return intList
}

func checkEven(data int) bool {
	if data%2 == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	result := 0

	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	aString := sc.Text()

	aStringList := strings.Split(aString, "")

	aIntList := convertStringListToIntList(aStringList)

	for _, v := range aIntList {
		if checkEven(v) {
			result++
		}
	}

	fmt.Println(result)
}
