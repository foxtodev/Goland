package fibonacci

//
func FibonacciCycle(number int) int {
	fib1, fib2 := 0, 1
	for ; number > 0; number-- {
		fib1, fib2 = fib2, fib1+fib2
	}
	return fib1
}

//
func FibonacciRecursion(number int) int {
	if number < 2 {
		return number
	}
	return FibonacciRecursion(number-1) + FibonacciRecursion(number-2)
}

//
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

//
func FibonacciFunc() func() int {
	fib1, fib2 := 1, 0
	return func() int {
		fib1, fib2 = fib2, fib1+fib2
		return fib1
	}
}

//
func FibonacciMap(number int) int {
	fmap := make(map[int]int)
	fmap[0], fmap[1] = 0, 1
	for i := 2; i <= number; i++ {
		fmap[i] = fmap[i-1] + fmap[i-2]
	}
	return fmap[number]
}

//
func FibonacciMapV2(number int) int {
	fib := map[int]int{
		0: 0,
		1: 1,
	}
	return FibonacciMapV2Rec(number, fib)
}

func FibonacciMapV2Rec(number int, fibc map[int]int) int {
	val, exists := fibc[number]
	if exists {
		return val
	}
	fibc[number] = FibonacciMapV2Rec(number-1, fibc) + FibonacciMapV2Rec(number-2, fibc)
	return fibc[number]
}

//
var fibcache = map[int]int{}

func FibonacciCache(number int) int {
	val, exists := fibcache[number]
	if !exists {
		val = FibonacciCacheRec(number)
		fibcache[number] = val
	}
	return val
}

func FibonacciCacheRec(number int) int {
	if number < 2 {
		return number
	}
	return FibonacciCacheRec(number-1) + FibonacciCacheRec(number-2)
}

//
func FibonacciMapFunc() func() int {
	fib := make(map[int]int)
	fib[0], fib[1] = 0, 1
	number := -1
	return func() int {
		number++
		if val, exists := fib[number]; exists {
			return val
		}
		fib[number] = fib[number-1] + fib[number-2]
		return fib[number]
	}
}
