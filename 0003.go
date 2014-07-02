package main

import "fmt"
import "math"

func isPrime(n int) bool {
	root := int(math.Sqrt(float64(n)))
	for i := 2; i < root; i++ {
		if n%2 == 0 {
			return true
		}
	}
	return false
}

func main() {
	const NUM int = 600851475143
	n := NUM
	max := 0

	for i := 2; i <= n; i++ {
		for !isPrime(i) && n%i == 0 {
			max = i
			n /= i
		}
	}

	fmt.Println(max)
}
