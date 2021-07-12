<br />
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
		0: 0,
		1: 1,
	}
	return FibonacciMapV2Rec(number, fib)
}

func FibonacciMapV2Rec(number int, fibc map[int]int) int {
	val, exists := fibc[number]
	if exists == true {
		return val
	}
	fibc[number] = FibonacciMapV2Rec(number-1, fibc) + FibonacciMapV2Rec(number-2, fibc)
	return fibc[number]
}
```
<br />

#### Calculating the fibonacci number by recursion with map ver2

```go
func FibonacciMapFunc() func() int {
	fib := make(map[int]int)
	fib[0], fib[1] = 0, 1
	number := -1
	return func() int {
		number++
		if val, exists := fib[number]; exists == true {
			return val
		}
		fib[number] = fib[number-1] + fib[number-2]
		return fib[number]
	}
}
```
<br />