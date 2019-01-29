package arrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTapeEquilibrium(t *testing.T) {
	testCases := []struct {
		A        []int
		expected int
	}{
		{
			A:        []int{3, 1, 2, 4, 3},
			expected: 1,
		},
		{
			A:        []int{-1000, 1000},
			expected: 2000,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TapeEquilibrium[%d]", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.expected, TapeEquilibrium(testCase.A))
		})
	}
}

func BenchmarkTapeEquilibrium(b *testing.B) {
	A := []int{3, 1, 2, 4, 3}
	for n := 0; n < b.N; n++ {
		TapeEquilibrium(A)
	}
}
