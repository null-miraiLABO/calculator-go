package main

import (
	"fmt"
	"strconv"
)

func main() {
	var total int
	limit := 10
	formula := make([]string, limit)

	for i := 0; i < limit; i++ {
		fmt.Scanf("%s", &formula[i])
		if formula[i] == "" {
			break
		}
	}

	fmt.Println(formula)

	tmp, _ := strconv.Atoi(formula[0])
	total = tmp

	for i := 0; i < len(formula); i++ {

		switch formula[i] {
		case "+":
			tmp, _ = strconv.Atoi(formula[i+1])
			total = total + tmp
		case "-":
			tmp, _ = strconv.Atoi(formula[i+1])
			total = total - tmp
		case "*":
			tmp, _ = strconv.Atoi(formula[i+1])
			total = total * tmp
		case "/":
			tmp, _ = strconv.Atoi(formula[i+1])
			total = total / tmp
		}

	}

	fmt.Printf("= %d \n", total)
}
