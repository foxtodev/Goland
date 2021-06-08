package main

import (
	"fmt"
	"math"
	"strings"
)

func PrimeSieve(num uint) []uint {
	var nums []uint
	arr := make([]bool, num)
	for i := uint(2); i <= uint(math.Sqrt(float64(num)+1)); i++ {
		if arr[i] == false {
			for j := i * i; j < num; j += i {
				arr[j] = true
			}
		}
	}
	for i, isComposite := range arr {
		if i > 1 && !isComposite {
			nums = append(nums, uint(i))
		}
	}
	return nums
}

func main() {
	var number uint
	fmt.Print("Please enter number ")
	fmt.Scanf("%d\n", &number)
	fmt.Printf("Prime numbers in the range [0..%d]: ", number)
	//fmt.Println(PrimeSieve(number))
	//fmt.Println(strings.Trim(fmt.Sprint(PrimeSieve(number)), "[]"))
	fmt.Println(strings.Trim(strings.Replace(fmt.Sprint(PrimeSieve(number)), " ", ", ", -1), "[]"))
}
