package gowther

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	ts := []struct {
		name     string
		input    int
		expected []interface{}
	}{
		{"test one", 1, []interface{}{1}},
		{"test two", 3, []interface{}{1, 2, "Fizz"}},
		{"test three", 5, []interface{}{1, 2, "Fizz", 4, "Buzz"}},
		{"test four", 15, []interface{}{1, 2, "Fizz", 4, "Buzz", "Fizz", 7, 8, "Fizz", "Buzz", 11, "Fizz", 13, 14, "Fizz Buzz"}},
	}

	for _, tc := range ts {
		t.Run(tc.name, func(t *testing.T) {
			act := FizzBuzz(tc.input)

			assert.NotNil(t, act)
			assert.Equal(t, tc.expected, act)
		})
	}
}

func BenchmarkFizzBuzz(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FizzBuzz(1000)
	}
}

func TestIsFizz(t *testing.T) {
	ts := []struct {
		name     string
		input    int
		expected bool
	}{
		{"test one", 1, false},
		{"test two", 2, false},
		{"test three", 3, true},
		{"test four", 4, false},
		{"test five", 5, false},
	}

	for _, tc := range ts {
		t.Run(tc.name, func(t *testing.T) {
			act := IsFizz(tc.input)
			assert.Equal(t, tc.expected, act)
		})
	}
}

func BenchmarkIsFizz(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsFizz(30)
	}
}

func TestIsBuzz(t *testing.T) {
	ts := []struct {
		name     string
		input    int
		expected bool
	}{
		{"test one", 1, false},
		{"test two", 2, false},
		{"test three", 3, false},
		{"test four", 4, false},
		{"test five", 5, true},
	}

	for _, tc := range ts {
		t.Run(tc.name, func(t *testing.T) {
			act := IsBuzz(tc.input)
			assert.Equal(t, tc.expected, act)
		})
	}
}

func BenchmarkIsBuzz(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsBuzz(50)
	}
}

func TestIsFizzBuzz(t *testing.T) {
	ts := []struct {
		name     string
		input    int
		expected bool
	}{
		{"test one", 10, false},
		{"test two", 20, false},
		{"test three", 30, true},
		{"test four", 40, false},
		{"test five", 50, false},
	}

	for _, tc := range ts {
		t.Run(tc.name, func(t *testing.T) {
			act := IsFizzBuzz(tc.input)
			assert.Equal(t, tc.expected, act)
		})
	}
}

func BenchmarkIsFizzBuzz(b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsFizzBuzz(30)
	}
}
