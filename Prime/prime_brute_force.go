package main

import (
	"fmt"
	"strings"
)

func PrimeBruteForce(num uint) []uint {
	var nums []uint
	var n uint
	for i := uint(2); i <= num; i++ {
		for n = 2; n <= i; n++ {
			if i%n == 0 {
				break
			}
		}
		if n == i {
			nums = append(nums, i)
		}
	}
	return nums
}

func main() {
	var number uint
	fmt.Print("Please enter number ")
	fmt.Scanf("%d\n", &number)
	fmt.Printf("Prime numbers in the range [0..%d]: ", number)
	//fmt.Println(PrimeBruteForce(number))
	//fmt.Println(strings.Trim(fmt.Sprint(PrimeBruteForce(number)), "[]"))
	fmt.Println(strings.Trim(strings.Replace(fmt.Sprint(PrimeBruteForce(number)), " ", ", ", -1), "[]"))
}
