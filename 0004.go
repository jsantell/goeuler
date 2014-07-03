package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isPalindrone(n int) bool {
	stringNum := strings.Split(strconv.Itoa(n), "")
	var rStringNum string

	for i := len(stringNum) - 1; i >= 0; i-- {
		rStringNum += stringNum[i]
	}

	return rStringNum == strings.Join(stringNum, "")
}

func main() {
	largest := 0
	for i := 999; i > 0; i-- {
		for j := i; j > 0; j-- {
			sum := i * j
			if isPalindrone(sum) && sum > largest {
				largest = sum
			}
		}
	}

	fmt.Println(largest)
}
