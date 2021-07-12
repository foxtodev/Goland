package main

import (
	"fmt"
	"strings"
)

func IsPrime(num uint, d uint) bool {
	if num < 2 {
		return false
	}
	if d*d > num || num == 2 {
		return true
	}
	if num%d == 0 {
		return false
	}
	return IsPrime(num, d+1)
}

func Prime(num uint) []uint {
	var nums []uint
	for i := uint(0); i <= num; i++ {
		if IsPrime(i, 2) {
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
	//fmt.Println(Prime(number))
	//fmt.Println(strings.Trim(fmt.Sprint(Prime(number)), "[]"))
	fmt.Println(strings.Trim(strings.Replace(fmt.Sprint(Prime(number)), " ", ", ", -1), "[]"))
}
