package main

import "fmt"

func main() {
	a := make([]int, 5)
	a[1] = 2
	a[2] = 3

	for i, v := range a {
		fmt.Println(i, v)
		// If you want only the values, you can skip the index using an underscore:
		// for _, v := range  {}

	}

	x := "hello"
	for _, c := range x {
		fmt.Printf("%c \n", c)
	}

	fmt.Println("MOD => ", 0%2)
}
