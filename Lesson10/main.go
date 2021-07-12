package main

import (
	"fmt"
	"test/fibonacci"
	"test/prime"
	"test/sort"
)

func main() {
	//Sort
	data := []int{7, 5, 2, 1, 4, 3, 8, 9, 0, 6}
	fmt.Println(sort.SortBubble(data))
	fmt.Println(sort.SortInsertion(data))
	fmt.Println(sort.SortQuick(data))
	fmt.Println(sort.SortGoSlice(data))

	//Fibonacci
	number := 20

	for i := 0; i <= number; i++ {
		fmt.Print(fibonacci.FibonacciCycle(i))
		fmt.Print(" ")
	}
	fmt.Println(" ")

	for i := 0; i <= number; i++ {
		fmt.Print(fibonacci.FibonacciRecursion(i))
		fmt.Print(" ")
	}
	fmt.Println(" ")

	for i := 0; i <= number; i++ {
		fmt.Print(fibonacci.FibonacciSlice(i))
		fmt.Print(" ")
	}
	fmt.Println(" ")

	for i := 0; i <= number; i++ {
		fmt.Print(fibonacci.FibonacciMapV2(i))
		fmt.Print(" ")
	}
	fmt.Println(" ")

	for i := 0; i <= number; i++ {
		fmt.Print(fibonacci.FibonacciCache(i))
		fmt.Print(" ")
	}
	fmt.Println(" ")

	f := fibonacci.FibonacciMapFunc()
	for i := 0; i <= number; i++ {
		fmt.Print(f())
		fmt.Print(" ")
	}
	fmt.Println(" ")

	//Prime
	var numberPrime uint = 20
	fmt.Println(prime.Prime(numberPrime))
	fmt.Println(prime.PrimeBruteForce(numberPrime))
	fmt.Println(prime.PrimeSieve(numberPrime))
}
