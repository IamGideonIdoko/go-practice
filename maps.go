package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["James"] = 42
	m["Amy"] = 24

	fmt.Println(m["James"])
	fmt.Println(m)

	// deleteing a value from the map
	delete(m, "Amy")
	fmt.Println(m)

	mp := map[int]int{
		8: 42,
		2: 6,
		4: 9,
		5: 3}
	fmt.Println(mp)
}
