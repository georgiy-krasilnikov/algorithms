package main

import (
	"fmt"

	"algo/algo"
)

func main() {
	sub := "lilila"
	s := "lililo lilila"
	res := algo.KMP(s, sub)
	fmt.Println(res)
}
