package main

import "fmt"

func pyramid(m int) {
	mm := m * 2
	for i := 0; i <= m; i++ {
		for j := 0; j < mm; j++ {
			fmt.Printf(" ")
		}
		mm -= 2
		n := i*2 - 1
		for k := 0; k < n; k++ {
			fmt.Printf("* ")
		}
		fmt.Printf("\n")
	}
}

func dimond(m int) {
	mm := m * 2
	for i := 0; i < m; i++ {
		for j := 0; j < mm; j++ {
			fmt.Printf(" ")
		}
		mm -= 2
		n := i*2 - 1
		for k := 0; k < n; k++ {
			fmt.Printf("* ")
		}
		fmt.Printf("\n")
	}
	mm += 4
	for i := m; i > 0; i-- {
		for j := 0; j < mm; j++ {
			fmt.Printf(" ")
		}
		mm += 2
		n := (i-2)*2 - 1
		for k := 0; k < n; k++ {
			fmt.Printf("* ")
		}
		fmt.Printf("\n")

	}
}

func main() {
	pyramid(15)
	dimond(15)

}
