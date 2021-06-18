package main

import (
	"fmt"
)

func FibonacciCycle(number int) int {
	fib1, fib2 := 0, 1
	for ; number > 0; number-- {
		fib1, fib2 = fib2, fib1+fib2
	}
	return fib1
}

func FibonacciRecursion(number int) int {
	if number < 2 {
		return number
	}
	return FibonacciRecursion(number-1) + FibonacciRecursion(number-2)
}

func FibonacciSlice(number int) int {
	if number < 2 {
		return number
	}
	fib := []int{0, 1}
	//for ; number > 0; number-- {
	for i := 2; i <= number; i++ {
		fib = []int{fib[1], fib[0] + fib[1]}
	}
	return fib[1]
}

func FibonacciFunc() func() int {
	fib1, fib2 := 1, 0
	return func() int {
		fib1, fib2 = fib2, fib1+fib2
		return fib1
	}
}

func FibonacciMap(number int) int {
	fmap := make(map[int]int)
	fmap[0], fmap[1] = 0, 1
	for i := 2; i <= number; i++ {
		fmap[i] = fmap[i-1] + fmap[i-2]
	}
	return fmap[number]
}

func FibonacciMapFunc() func() int {
	fib := make(map[int]int)
	number := 0
	return func() int {
		fib[0], fib[1] = 0, 1
		number++
		if number < 2 {
			return fib[0]
		}
		fib[number-1] = fib[number-2] + fib[number-3]
		return fib[number-1]
	}
}

func main() {

	var number int

	fmt.Print("Enter number: ")
	fmt.Scanln(&number)

	f := FibonacciMapFunc()
	for i := 0; i <= number; i++ {
		fmt.Print(f())
		fmt.Print(" ")
	}
	fmt.Println(" ")

}
