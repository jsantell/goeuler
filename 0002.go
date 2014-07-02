package main

import "fmt"

func fibonacci(max int) chan int {
	var z int
	x := 0
	y := 1
	ch := make(chan int)

	go func() {
		for y+x < max {
			z = x + y
			x = y
			y = z
			ch <- y
		}
		close(ch)
	}()

	return ch
}

func main() {
	const MAX = 4000000

	total := 0
	for x := range fibonacci(MAX) {
		fmt.Println(x, x%2 == 0, total)
		if x%2 == 0 {
			total += x
		}
	}
	fmt.Println(total)
}
