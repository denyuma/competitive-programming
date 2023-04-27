package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)

	for i := 8; i > 0; i-- {
		sc.Scan()
		aString := sc.Text()
		aStringList := strings.Split(aString, "")
		for j := 0; j < 8; j++ {
			if aStringList[j] == "*" {
				switch j {
				case 0:
					fmt.Printf("a%v\n", i)
				case 1:
					fmt.Printf("b%v\n", i)
				case 2:
					fmt.Printf("c%v\n", i)
				case 3:
					fmt.Printf("d%v\n", i)
				case 4:
					fmt.Printf("e%v\n", i)
				case 5:
					fmt.Printf("f%v\n", i)
				case 6:
					fmt.Printf("g%v\n", i)
				case 7:
					fmt.Printf("h%v\n", i)
				}
			}
		}
	}
}
