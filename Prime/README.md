## Finding prime numbers

### number = 1000

```
go test -bench . -benchmem  prime_test.go
```

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz
BenchmarkPrime-2                   16244             69427 ns/op            4088 B/op          9 allocs/op
BenchmarkPrimeSieve-2             446959              2856 ns/op            5112 B/op         10 allocs/op
BenchmarkPrimeBruteForce-2       1202352               987.3 ns/op             8 B/op          1 allocs/op
PASS
ok      command-line-arguments  5.400s
```
## Naive brute force method win))))