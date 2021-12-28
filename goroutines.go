package main

import (
	"fmt"
	"time"
)

func out(from, to int) {
	for i := from; i <= to; i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Println(i)
	}
}
func main() {
	/* out(0, 5)
	out(6, 10) */
	go out(0, 5) // the go keyword creates a go routine
	go out(6, 10)

	/*
		Our program has 3 virtual threads: the 2 function calls, and main(). Our 2 function calls get executed concurrently, and main() does not wait for them to finish.
	*/

	time.Sleep(500 * time.Millisecond) /* make main function not to terminate until after some time  */
}
