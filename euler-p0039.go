package main

import (
	"fmt"
//	"sandbox/euler"
)

const p int = 121

type triple struct {
	m int
	n int
	a int
	b int
	c int
}

func main() {
		
	var sums [p]int
	var triples [p]([]triple)
	
	for i := 0; i < p; i++ { triples[i] = make([]triple, 0) }

	for n := 1; n < p/2; n++ {
		for m := n + 1; 2*m*(m+n) <= p; m++ {
			for k := 1; ; k++ {
				sum := k*2*m*(m+n)
				//fmt.Printf("%d %d = %d\n", n, m, sum)
				if sum > p { break }
				sums[sum]++
				triples[sum] = append(triples[sum], 
					triple{m, n, k*((m*m)-(n*n)),k*2*m*n,k*((m*m)+(n*n))})
			}
		}
	}
	fmt.Printf("%d\n", sums[120]);
	for _, t := range triples[120] { fmt.Printf("[%d %d] %d %d %d\n", t.m, t.n, t.a, t.b, t.c) }
}
