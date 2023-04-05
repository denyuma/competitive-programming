package main

import "fmt"

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	var S string
	var divide = [4]string{"dream", "dreamer", "erase", "eraser"}
	for i, v := range divide {
		divide[i] = reverse(v)
	}
	can := true
	fmt.Scan(&S)
	S = reverse(S)

	for can {
		can2 := false
		if len(S) == 0 {
			break
		}
		for j := 0; j < len(divide); j++ {
			d := divide[j]
			if len(d) > len(S) {
				continue
			}
			if S[:len(d)] == d {
				can2 = true
				S = S[len(d):]
				break
			}
		}

		if !can2 {
			can = false
			break
		}
	}

	if can {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
