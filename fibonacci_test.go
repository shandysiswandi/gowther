package gowther

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFibonacci(t *testing.T) {
	ts := []struct {
		name     string
		input    int
		expected []int
	}{
		{"test one", 1, []int{0, 1}},
		{"test two", 3, []int{0, 1, 1}},
		{"test three", 5, []int{0, 1, 1, 2, 3}},
		{"test four", 15, []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377}},
	}

	for _, tc := range ts {
		t.Run(tc.name, func(t *testing.T) {
			act := Fibonacci(tc.input)

			assert.NotNil(t, act)
			assert.Equal(t, tc.expected, act)
		})
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fibonacci(1000)
	}
}
