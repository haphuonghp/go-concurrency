package main

import (
	"fmt"
	"time"
)

func fun(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Direct call
	fun("direct call")

	// TODO: write goroutine with different variants for function call.

	// goroutine function call
	go fun("Alo 1	")
	// goroutine with anonymous function
	go func() {
		fun("Alo 2")
	}()
	// goroutine with function value call
	export := fun
	go export("Alo 3")

	// wait for goroutines to end
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("done..")
}
