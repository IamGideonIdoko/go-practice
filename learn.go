package main

import "fmt"

func mars_age(earth_age int) int {
	return (earth_age * 365) / 687
}

func main() {
	/* var age int
	fmt.Printf("Enter age: ")
	fmt.Scanln(&age)

	mars := mars_age(age)
	fmt.Println(mars) */
	/*
		x := 42
		p := &x */

	type Card struct {
		name string
		age  int
	}

	fmt.Printf("%T", Card{name: "James", age: 24})
	fmt.Println(Card{name: "James", age: 24})
}
