/*
Variadic functions are functions that can be called with any number of arguments.
Eg. fmt.Println("a", "b", "c")
*/
package main

import "fmt"

func sum(nums ...int) { // ... defines any number of arg, like the rest operator in javasscript
	fmt.Println("Nums => ", nums)
	total := 0
	for _, v := range nums {
		total += v
	}
	fmt.Println(total)
}
func main() {
	sum(2, 4, 6)
	sum(42, 8)
	sum(1, 2, 3, 4, 5, 6)

	s := []int{42, 33, 22, 11}
	sum(s...)
}
