package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Print("Please enter a number ")
	fmt.Scanf("%d\n", &number)
	fmt.Printf("Hundreds: %d, dozens: %d, units: %d\n", number/100%10, number/10%10, number%10)
}
