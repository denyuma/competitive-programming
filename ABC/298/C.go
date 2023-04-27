// WA

package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func include(slice []int, target int) bool {
	for _, num := range slice {
		if num == target {
			return true
		}
	}

	return false
}

func main() {
	var n, q int
	boxies := make([][]int, n)
	for i := 0; i < n; i++ {
		boxies[i] = make([]int, q)
	}

	fmt.Println(reflect.TypeOf(boxies[0]))
	var sc = bufio.NewScanner(os.Stdin)

	fmt.Scan(&n)
	fmt.Scan(&q)

	for i := 0; i < q; i++ {
		sc.Scan()
		aString := sc.Text()
		aStringArray := strings.Split(aString, " ")
		aIntArray := make([]int, q)

		for i, str := range aStringArray {
			aIntArray[i], _ = strconv.Atoi(str)
		}
		queryNumber := aIntArray[0]
		fmt.Println("queryNumber", queryNumber)

		switch queryNumber {
		case 1:
			cardNumber := aIntArray[1]
			boxNumber := aIntArray[2]

			boxies[boxNumber] = append(boxies[boxNumber], cardNumber)
		case 2:
			boxNumber := aIntArray[1]
			sort.Ints(boxies[boxNumber])
			for _, v := range boxies[boxNumber] {
				fmt.Printf("%d ", v)
			}
			fmt.Printf("\n")

		case 3:
			cardNumber := aIntArray[1]
			for j, eachBox := range boxies {
				if include(eachBox, cardNumber) {
					fmt.Printf("%v ", j+1)
				}
			}
			fmt.Printf("\n")

		}

	}
}
