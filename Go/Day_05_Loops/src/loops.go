package main

import "fmt"

func main() {
	var T int
	var a, b, N uint32

	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&a, &b, &N)

		sum := a

		for j := uint32(0); j < N; j++ {
			sum += (1 << j) * b
			fmt.Printf("%d ", sum)
		}
		fmt.Println()
	}
}
