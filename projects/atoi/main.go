package main

import (
	"fmt"
)

func main() {
	var st, result string
	fmt.Scan(&st)

	for _, c := range st {
		num := int(c - '0')
		result += fmt.Sprintf("%d", num*num)
	}

	fmt.Println(result)
}
