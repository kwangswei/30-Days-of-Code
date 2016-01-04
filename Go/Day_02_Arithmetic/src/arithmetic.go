package main

import (
	"fmt"
)

func main() {
	var m float32
	var t, x int

	fmt.Scanf("%f %d %d", &m, &t, &x)

	var tip = (m * float32(t)) / 100
	var tax = (m * float32(x)) / 100

	fmt.Printf("The final price of the meal is $%v.\n", int(m+tip+tax+0.5))
}
