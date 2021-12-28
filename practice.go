package main

import "fmt"

func main() {
	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}

	var input string
	fmt.Scanln(&input)

	res := append(results, input)

	output := 0

	for _, v := range res {
		if v == "w" {
			output += 3
		} else if v == "d" {
			output += 1
		}
	}

	fmt.Println(output)

}
