package main

import "fmt"

func strip(bs []byte) []byte {
	var i int
	for i = 0; bs[i] == 0; i++ { }	
	return bs[i:]
}

func multiplication(m, n []byte) []byte {
	res := make([]byte, 0)
	
	var lines [][]byte	

	lines = make([][]byte, len(n))
	var transfer byte 
	transfer = 0
	for i := len(n) - 1; i >= 0; i-- {
		lines[i] = make([]byte, 0)
		for k := i; k < len(n) - 1; k++ { lines[i] = append(lines[i], 0) }
		for j := len(m) - 1; j >= 0; j-- {
			x := m[j] * n[i] + transfer			
			transfer = byte(x / 10)
			remainder := x % 10
			lines[i] = append([]byte{remainder}, lines[i]...)

			if j == 0 { 
				lines[i] = append([]byte{transfer}, lines[i]...)
			}
		}
		for k := i; k >= 0; k-- { lines[i] = append([]byte{0}, lines[i]...) }
	}

	transfer = 0
	for i := len(lines[0]) - 1; i >= 0; i-- {
		var c byte
		c = transfer
		for j := 0; j < len(n); j++ {
			c += lines[j][i] 
		}
		transfer = byte(c / 10)
		remainder := c % 10
		res = append([]byte{remainder}, res...)

		if i == 0 { res = append([]byte{transfer}, res...) }
	}

	return res
}       

func diffuse(x int) []byte {
	res := make([]byte, 0)
	y := x
	res = append(res, byte(y / 100))
	y %= 100
	res = append(res, byte(y / 10))
	y %= 10
	res = append(res, byte(y))
	return res
}

func main() {
	for a := 2; a <= 100; a++ {		
		ab := diffuse(a)
		res := ab
		for b := 2; b <= 100; b++ {
			res = multiplication(res, ab)
			for _, c := range strip(res) { fmt.Printf("%d", c) }
			fmt.Printf("\n")
		}
	}
}
