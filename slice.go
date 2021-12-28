package main

import "fmt"

func main() {
	a := [5]int{0, 2, 4, 6, 8}
	var s []int = a[1:3]

	fmt.Println(s)
	fmt.Println(s[1])

	// making slice in go
	b := make([]int, 5)
	b = append(b, 8)
	fmt.Println(b)
}
