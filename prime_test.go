package gowther

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrimes(t *testing.T) {
	ts := []struct {
		name     string
		input    int
		expected []int
	}{
		{"test one", 1, []int{}},
		{"test two", 3, []int{2, 3}},
		{"test three", 5, []int{2, 3, 5}},
		{"test four", 15, []int{2, 3, 5, 7, 11, 13}},
	}

	for _, tc := range ts {
		t.Run(tc.name, func(t *testing.T) {
			act := Primes(tc.input)

			assert.NotNil(t, act)
			assert.Equal(t, tc.expected, act)
		})
	}
}

func BenchmarkPrimes(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Primes(1000)
	}
}

func TestIsPrime(t *testing.T) {
	ts := []struct {
		name     string
		input    int
		expected bool
	}{
		{"test one", 1, false},
		{"test two", 2, true},
		{"test three", 3, true},
		{"test four", 4, false},
	}

	for _, tc := range ts {
		t.Run(tc.name, func(t *testing.T) {
			act := IsPrime(tc.input)
			assert.Equal(t, tc.expected, act)
		})
	}
}

func BenchmarkPrime(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsPrime(100)
	}
}
