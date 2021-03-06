package main

import (
	"math"
	"testing"
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

func BenchmarkPrime(b *testing.B) {
	var number uint = 100000
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Prime(number)
	}
}

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

func BenchmarkPrimeSieve(b *testing.B) {
	var number uint = 100000
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		PrimeSieve(number)
	}
}

func PrimeBruteForce(num uint) []uint {
	var nums []uint
	var n uint
	for i := uint(2); i <= num; i++ {
		for n = 2; n <= i; i++ {
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

func BenchmarkPrimeBruteForce(b *testing.B) {
	var number uint = 100000
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		PrimeBruteForce(number)
	}
}

func main() {

}
