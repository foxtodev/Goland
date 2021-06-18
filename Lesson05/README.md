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