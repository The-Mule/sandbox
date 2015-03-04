package main

import "fmt"

const p int = 1001

type Triple struct {
	a int
	b int
	c int	
}

type Triples []Triple

func (t *Triple) normalize() {
	if t.a > t.b { t.a, t.b = t.b, t.a }
	if t.a > t.c { t.a, t.c = t.c, t.a }
	if t.b > t.c { t.b, t.c = t.c, t.b }
}

func (ts *Triples) contains(t Triple) bool {
	for _, x := range *ts {
		if t.a == x.a && t.b == x.b && t.b == x.b { return true } 
	}
	return false
}

func main() {
		
	var sums [p]int
	var triples [p]Triples
	
	for i := 0; i < p; i++ { triples[i] = make(Triples, 0) }

	max := 0
	for n := 1; n < p/2; n++ {
		for m := n + 1; 2*m*(m+n) <= p; m++ {
			for k := 1; ; k++ {
				sum := k*2*m*(m+n)
				if sum > p { break }
				t := Triple{k*((m*m)-(n*n)),k*2*m*n,k*((m*m)+(n*n))}
				t.normalize()
				if !triples[sum].contains(t) {
					sums[sum]++
					triples[sum] = append(triples[sum], t)
					if sums[sum] > sums[max] { max = sum }
				}
			}
		}
	}
	for _, t := range triples[max] { fmt.Printf("%d %d %d\n", t.a, t.b, t.c) }
	fmt.Printf("%d (%d)\n", sums[max], max)
}
