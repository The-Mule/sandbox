package main

import "fmt"

type fraction struct {
	a int
	b int
}

func gcd(a, b int) int {
	if b == 0 { return a }
	return gcd(b, a % b)
}

func simplify(f fraction) fraction {
	d := gcd(f.a, f.b)
	return fraction{f.a/d, f.b/d}
}

func equal(f1, f2 fraction) bool {
	f1_s := simplify(f1)
	f2_s := simplify(f2)
	return f1_s.a == f2_s.a && f1_s.b == f2_s.b
}

func shared_digits(f fraction) []byte {
	res := make([]byte, 0)
	a1 := byte(f.a / 10)
	a2 := byte(f.a % 10)
	b1 := byte(f.b / 10)
	b2 := byte(f.b % 10)

	if a1 == b1 || a1 == b2 { res = append(res, a1) }
	if a1 != a2 && (a2 == b1 || a2 == b2) { res = append(res, a2) }

	return res
}

func remove_digit(f fraction, d byte) fraction {
	var a,b int
	if byte(f.a / 10) == d { a = f.a % 10 }
	if byte(f.a % 10) == d { a = f.a / 10 } 

	if byte(f.b / 10) == d { b = f.b % 10 }
	if byte(f.b % 10) == d { b = f.b / 10 }

	return fraction{a,b}
}

func main() {
	res := make([]fraction, 0)
	for i := 10; i < 99; i++ {
		for j := i + 1; j < 99; j++ {		
			if i % 10 == 00 && j % 10 == 0 { continue }
			f := fraction{i,j}
			sd := shared_digits(f)
			for k := 0; k < len(sd); k++ {
				g := remove_digit(f, sd[k])
				if equal(f, g) { res = append(res, g) }
			}			
		}
	}
	a := 1
	b := 1
	for _, f := range res {
		a *= f.a
		b *= f.b
	}

	solution := simplify(fraction{a,b})

	fmt.Printf("%d/%d\n", solution.a, solution.b)
}
