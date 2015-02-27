package main

import "fmt"
import "sandbox/euler"

func compute(lb, up int) {
	for i := lb; i < up; i++ {
		s := euler.Slice(i)		
		hit := true
		for j := 0; j < len(s); j++ {
			r := euler.Rotate(s,j)
			if !euler.PrimalityTest(euler.Number(r)) {
				hit = false
				break			
			}			
		}
		if hit { fmt.Printf("%d\n", i) }
	}
}

func main() {

	compute(2, 1000000)
}
