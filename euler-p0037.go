package main

import (
	"fmt"
	"sandbox/euler"
)


func check(n int) bool {
	s := euler.Slice(n, 10)	
	for i := 0; i < len(s); i++ { 
		if !euler.PrimalityTest(euler.Number(s[i:], 10)) { return false }
		if !euler.PrimalityTest(euler.Number(s[:i+1], 10)) { return false }


	}
	return true
}

func main() {
	sum := 0
	for i := 10; i < 1000000; i++ { 
		if check(i) {
			fmt.Printf("%d\n", i)
			sum += i
		}
	}
	fmt.Printf("%d\n", sum)
}
