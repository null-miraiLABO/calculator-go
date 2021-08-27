package main

import (
	"fmt"
	"strconv"
)

func main() {
	var cv int
	limit := 10
	formula := make([]string, limit)

	for i := 0; i < limit; i++ {
		fmt.Scanf("%s", &formula[i])
		if formula[i] == "" {
			break
		}
	}

	fmt.Println(formula)

	for i := 0; i < len(formula); i++ {
		tmp := 0

		switch formula[i] {
		case "*":
			fmt.Println("*")
		case "/":
			fmt.Println("/")
		case "+":
			fmt.Println("+")
		case "-":
			//fmt.Println(minus(formula[i-1],formula[i+1]))
		default:
			tmp, _ = strconv.Atoi(formula[i])
			cv += tmp
		}

	}

	fmt.Println(cv)
}

/*
func minus(x, y int) int {
	return x - y
}
*/
