package fibonacci

import "testing"

var table = []struct {
	arg  int
	want int
}{
	{0, 0},
	{1, 1},
	{10, 55},
	{20, 6765},
}

const N = 20

var sink int

func TestFibonacciCycle(t *testing.T) {
	for _, entry := range table {
		got := FibonacciCycle(entry.arg)
		if got != entry.want {
			t.Errorf("For %d got %d want %d", entry.arg, got, entry.want)
		}
	}
}

func BenchmarkFibonacciCycle(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		res = FibonacciCycle(N)
	}
	sink = res
}

func TestFibonacciRecursion(t *testing.T) {
	for _, entry := range table {
		got := FibonacciRecursion(entry.arg)
		if got != entry.want {
			t.Errorf("For %d got %d want %d", entry.arg, got, entry.want)
		}
	}
}

func BenchmarkFibonacciRecursion(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		res = FibonacciRecursion(N)
	}
	sink = res
}

func TestFibonacciSlice(t *testing.T) {
	for _, entry := range table {
		got := FibonacciSlice(entry.arg)
		if got != entry.want {
			t.Errorf("For %d got %d want %d", entry.arg, got, entry.want)
		}
	}
}

func BenchmarkFibonacciSlice(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		res = FibonacciSlice(N)
	}
	sink = res
}

func TestFibonacciMap(t *testing.T) {
	for _, entry := range table {
		got := FibonacciMap(entry.arg)
		if got != entry.want {
			t.Errorf("For %d got %d want %d", entry.arg, got, entry.want)
		}
	}
}

func BenchmarkFibonacciMap(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		res = FibonacciMap(N)
	}
	sink = res
}

func TestFibonacciMapV2(t *testing.T) {
	for _, entry := range table {
		got := FibonacciMapV2(entry.arg)
		if got != entry.want {
			t.Errorf("For %d got %d want %d", entry.arg, got, entry.want)
		}
	}
}

func BenchmarkFibonacciMapV2(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		res = FibonacciMapV2(N)
	}
	sink = res
}

func TestFibonacciCache(t *testing.T) {
	for _, entry := range table {
		got := FibonacciCache(entry.arg)
		if got != entry.want {
			t.Errorf("For %d got %d want %d", entry.arg, got, entry.want)
		}

		gotCache := FibonacciCacheRec(entry.arg)
		if got != entry.want {
			t.Errorf("For %d got %d want %d", entry.arg, gotCache, entry.want)
		}
	}
}

func BenchmarkFibonacciCache(b *testing.B) {
	var res int
	for i := 0; i < b.N; i++ {
		res = FibonacciCache(N)
	}
	sink = res
}
