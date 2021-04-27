package gowther

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	assert.Equal(t, "0.1.0", Version())
}

func BenchmarkVersion(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Version()
	}
}
