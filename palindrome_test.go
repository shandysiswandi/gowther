package gowther

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	ts := []struct {
		name     string
		input    string
		expected bool
	}{
		{"One Falsy", "", true},
		{"Two Falsy", "a", true},
		{"Three Falsy", "aa", true},
		{"Four Falsy", "aba", true},
		{"Five Falsy", "katak", true},

		{"One Truly", "-1", false},
		{"Two Truly", "asi", false},
		{"Three Truly", "ulie", false},
		{"Four Truly", "mamas", false},
		{"Five Truly", "okiy123", false},
	}

	for _, tc := range ts {
		t.Run(tc.name, func(t *testing.T) {
			act := IsPalindrome(tc.input)
			assert.Equal(t, tc.expected, act)
		})
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("1234567890abba0987654321")
	}
}
