// In Golang, routines can be made by adding "go" keyword in the starting of any function or method.

// Below written code is an example of sequential programming in Golang.

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	sum := 0
	for i := 0; i <= 10; i++ {
		time.Sleep(1 * time.Second)
		sum += i
		fmt.Println(sum)
	}
	eslaped := time.Since(start)
	fmt.Printf("This program has taken %v seconds", eslaped)
}
