package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func divideByTwo(data []int) []int {
	dividedData := make([]int, len(data))
	for i, v := range data {
		dividedData[i] = v / 2
	}
	return dividedData
}

func checkAllEven(data []int) bool {
	flag := true
	for _, v := range data {
		if v%2 != 0 {
			flag = false
			break
		}
	}
	return flag
}

func main() {
	var n int
	var result int
	var sc = bufio.NewScanner(os.Stdin)

	fmt.Scan(&n)

	sc.Scan()
	text := sc.Text()
	strList := strings.Split(text, " ")
	intList := make([]int, len(strList))
	for i, v := range strList {
		intList[i], _ = strconv.Atoi(v)
	}

	for {
		if checkAllEven(intList) {
			intList = divideByTwo(intList)
			result++
		} else {
			break
		}
	}

	fmt.Println(result)
}
