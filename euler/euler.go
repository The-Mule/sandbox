package euler

var primes []int = make([]int, 0)

const max int = 1000000

func PrimalityTest(n int) bool {
	for _, p := range primes { if n > p && n % p == 0 { return false } }
	for i := 5; i*i <= n; i += 6 {
		if n % i == 0 || n % (i+2) == 0 { return false }
	}
	primes = append(primes, n)	
	return true
}

func Pow(a, b int) int {
	var res int = 1
	for i := 0; i < b; i++ {
		res *= a
	}
	return res
}

func Number(s []byte) int {
	res := 0
	for i := len(s) - 1; i >= 0; i-- {
		res += int(s[len(s) - 1 - i]) * Pow(10, i)
	}
	return res
}

func Slice(n int) []byte {
	s := make([]byte, 0)
	started := false
	for i := 6; i >= 0; i-- {
		p := Pow(10, i)
		v := byte(n / p)
		n %= p
		if v > 0 && !started { started = true }
		if started { s = append(s, v) }		
	}
	return s
}

func Rotate(s []byte, l int) []byte {
	res := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = s[(i+l)%len(s)]
	}
	return res
}

