#### Calculating the fibonacci number by recursion

```go
func FibonacciRecursion(number int) int {
	if number < 2 {
		return number
	}
	return FibonacciRecursion(number-1) + FibonacciRecursion(number-2)
}
```
<br />

#### Calculating the fibonacci number by recursion with map

```go
func FibonacciMapV2(number int) int {
	fib := map[int]int{
		1: 0,
		2: 1,
	}
	return FibonacciMapV2Rec(number, fib)
}

func FibonacciMapV2Rec(number int, fcache map[int]int) int {
	if number < 2 {
		return number
	}
	fcache[number] = FibonacciMapV2Rec(number-1, fcache) + FibonacciMapV2Rec(number-2, fcache)
	return fcache[number]
}
```
<br />

#### Calculating the fibonacci number by recursion with map ver2

```go
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
```
<br />