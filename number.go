package main

import (
	"fmt"
)

func NumberStairs(n int) {

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Printf("\n")

	}

}

func main() {
	NumberStairs(9)

}
