package prime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const N = 20

var sink []uint

var want = []uint{2, 3, 5, 7, 11, 13, 17, 19}

func TestPrime(t *testing.T) {
	assert.Equal(t, Prime(N), want, "Prime() failed")
}

func BenchmarkPrime(b *testing.B) {
	var res []uint = Prime(N)
	sink = res
}

func TestPrimeSieve(t *testing.T) {
	assert.Equal(t, PrimeSieve(N), want, "Prime() failed")
}

func BenchmarkPrimeSieveb(t *testing.B) {
	var res []uint = PrimeSieve(N)
	sink = res
}

func TestPrimeBruteForce(t *testing.T) {
	assert.Equal(t, PrimeBruteForce(N), want, "Prime() failed")
}

func BenchmarkPrimeBruteForce(t *testing.B) {
	var res []uint = PrimeBruteForce(N)
	sink = res
}
