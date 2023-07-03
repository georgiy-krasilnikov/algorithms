package main

import (
	"algo/algo"
	"fmt"
)

func main() {
	// sub := "lilila"
	// s := "lililo lilila"
	//res := algo.KMP(s, sub)
	//fmt.Println(res)
	mas := algo.NewSet[string](4)
	r := mas.Add("hello")
	fmt.Println(r)
}
