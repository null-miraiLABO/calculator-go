package main

import (
	"fmt"
	"strconv"
)

func main() {
	formula := set()
	fmt.Println(formula)

	// 入力チェック後に出力
	if formula[0] == "" || formula[0] == "+" || formula[0] == "-" || formula[0] == "*" || formula[0] == "/" {
		fmt.Println("先頭には数字を入力して下さい。")
	} else {
		total := calc(formula)
		fmt.Printf("= %d \n", total)
	}

}

// 式をセット
func set() []string {
	limit := 10
	tmp := make([]string, limit)
	for i := 0; i < limit; i++ {
		fmt.Scanf("%s", &tmp[i])
		if tmp[i] == "" {
			break
		}
	}
	return tmp
}

// 計算する
func calc(formula []string) int {
	tmp, _ := strconv.Atoi(formula[0])
	total := tmp

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
	return total
}
