package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	split := strings.Split(s, "")
	join := strings.Join(split, "*")
	fmt.Println(join)
}
