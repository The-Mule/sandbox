package main

import (
	"fmt"
	"sandbox/euler"
)

func main() {

	sum := 0
	for i := 0; i < 1000000; i++ {
		slice := euler.Slice(i, 2)
		if euler.Palindromic(slice) {		
			slice = euler.Slice(i, 10)			
			if euler.Palindromic(slice) { sum += i }
		}
	}
	fmt.Printf("%d\n", sum)
}
