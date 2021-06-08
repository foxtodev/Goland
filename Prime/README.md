## Finding prime numbers

```
go test -bench . -benchmem  prime_test.go
```

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkPrime-2                  874538              1477 ns/op             248 B/op          5 allocs/op
BenchmarkPrimeSieve-2           10911757                93.79 ns/op            8 B/op          1 allocs/op
BenchmarkPrimeBruteForce-2      11889464                87.59 ns/op            8 B/op          1 allocs/op
PASS
ok      command-line-arguments  4.509s
```
## Naive brute force method win))))