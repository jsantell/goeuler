package main

import (
	"fmt"
	"math"
)

func main() {
	var sumOfSquares float64
	var squareOfSums float64

	for i := 1; i <= 100; i++ {
		sumOfSquares += math.Pow(float64(i), 2)
		squareOfSums += float64(i)
	}
	squareOfSums = math.Pow(squareOfSums, 2)

	fmt.Println(int(squareOfSums - sumOfSquares))
}
