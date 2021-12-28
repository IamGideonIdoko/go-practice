package main

import "fmt"

func main() {
	fmt.Println("Welcome to my quiz game!")

	var name string
	fmt.Printf("Enter your name: ")
	fmt.Scan(&name) /* Scan needs the memory location of the variable that's the reason for & */

	fmt.Println("Hello %v, welcome to the game!", name)
	fmt.Printf("%T\n", name)


	/* fmt.Printf("\nEnter your age: ")

	var age uint
	fmt.Scan(&age) 
	

	if age >= 10 {
		fmt.Println("Yay, you can play!")
	} */
}