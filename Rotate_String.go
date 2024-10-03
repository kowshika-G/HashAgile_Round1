package main

import (
	"fmt"
)

func main() {
	x := "abcdef"
	y := 4
	z := len(x)
	y = y % z

	r := ""

	for i := y; i < z; i++ {
		r += string(x[i])
	}
	for i := 0; i < y; i++ {
		r += string(x[i])
	}

	fmt.Println(r)
}
