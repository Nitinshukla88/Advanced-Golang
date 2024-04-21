package main

import (
	"fmt"
	"time"
)

func calcsquare(i int) {
	time.Sleep(1 * time.Second)
	fmt.Println(i * i)
}

func main() {
	start := time.Now()
	for i := 0; i <= 10000; i++ {
		go calcsquare(i)
	}
	eslaped := time.Since(start)
	time.Sleep(2 * time.Second)
	fmt.Printf("Our function has taken %v time", eslaped)
}
