
#### Sort
###### Used Test and Example

``
go test -v
``

```
=== RUN   TestSortBubble
--- PASS: TestSortBubble (0.00s)
=== RUN   TestSortInsertion
--- PASS: TestSortInsertion (0.00s)
=== RUN   TestSortQuick
--- PASS: TestSortQuick (0.00s)
=== RUN   TestSortGoSlice
--- PASS: TestSortGoSlice (0.00s)
=== RUN   ExampleSortBubble
--- PASS: ExampleSortBubble (0.00s)
PASS
ok      test/sort       0.004s
```

#### Prime
###### Used Testify and Benchmark

``
go test -v
``

```
=== RUN   TestPrime
--- PASS: TestPrime (0.00s)
=== RUN   TestPrimeSieve
--- PASS: TestPrimeSieve (0.00s)
=== RUN   TestPrimeBruteForce
--- PASS: TestPrimeBruteForce (0.00s)
PASS
ok      test/prime      0.004s
```

``
go test -bench=.
``

```
goos: linux
goarch: amd64
pkg: test/prime
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkPrime-2                1000000000               0.0000064 ns/op
BenchmarkPrimeSieveb-2          1000000000               0.0000020 ns/op
BenchmarkPrimeBruteForce-2      1000000000               0.0000042 ns/op
PASS
ok      test/prime      0.035s
```

#### Prime
###### Used Test and Benchmark

``
go test -v
``

```
=== RUN   TestFibonacciCycle
--- PASS: TestFibonacciCycle (0.00s)
=== RUN   TestFibonacciRecursion
--- PASS: TestFibonacciRecursion (0.00s)
=== RUN   TestFibonacciSlice
--- PASS: TestFibonacciSlice (0.00s)
=== RUN   TestFibonacciMap
--- PASS: TestFibonacciMap (0.00s)
=== RUN   TestFibonacciMapV2
--- PASS: TestFibonacciMapV2 (0.00s)
=== RUN   TestFibonacciCache
--- PASS: TestFibonacciCache (0.00s)
PASS
ok      test/fibonacci  0.002s
```

``
go test -bench=.
``

```
goos: linux
goarch: amd64
pkg: test/fibonacci
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkFibonacciCycle-2		145095332			8.302 ns/op
BenchmarkFibonacciRecursion-2		    27298			42852 ns/op
BenchmarkFibonacciSlice-2		  2764297			413.5 ns/op
BenchmarkFibonacciMap-2			   658933			 1804 ns/op
BenchmarkFibonacciMapV2-2		   562284			 2098 ns/op
BenchmarkFibonacciCache-2		173800128			6.794 ns/op
PASS
ok      test/fibonacci  11.179s
```

<br />
