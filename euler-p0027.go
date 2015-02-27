package main

import "fmt"

var primes []int

func primality_test(n int) bool {

	x := n
	if n < 0 { x = -x }

//	last_prime := 2

//	for i, p := range primes {
//		fmt.Printf("  checking %d\n", p)
	// 	if p == x { return true }
	// 	if p <  x && x % p == 0 { return false }
	// 	if p >  x && i > 0 { last_prime = primes[i-1] }
	// }
	
//	if last_prime < x {
		for i := 2; i < x; i++ {
//			fmt.Printf("  checking %d\n", i)
			if x % i == 0 { return false }
 		}
//		primes = append(primes, x)
//	}
	return true
}

func consecutive_primes(a int, b int) int {
	var i int
	for i = 0; primality_test(i*i + a*i + b); i++ {}
	return i
}

func main() {

	primes = make([]int, 0)
	primes = append(primes, 2, 3, 5, 7, 11, 13)	
	max_a := 1
	max_b := 1
	max_v := 1

	ch := make(chan bool)
	for i := -999; i < 1000; i++ {
		go func(a int) {
			for b := -999; b < 1000; b++ {
				count := consecutive_primes(a, b)
				if count > max_v { 
					max_a = a
					max_b = b
					max_v = count
				}
			}
			//fmt.Printf("%d finished\n", a)
			ch <- true
		}(i)
	}
	for i := -999; i < 1000; i++ { <- ch }
	fmt.Printf("%d %d - %d (%d)\n", max_a, max_b, max_v, max_a * max_b)		
}
