package prefixsums

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinAvgTwoSlice(t *testing.T) {
	testCases := []struct {
		A        []int
		expected int
	}{
		{
			A:        []int{4, 2, 2, 5, 1, 5, 8},
			expected: 1,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("MinAvgTwoSlice[%d]", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.expected, MinAvgTwoSlice(testCase.A))
		})
	}
}

func BenchmarkMinAvgTwoSlice(b *testing.B) {
	A := []int{4, 2, 2, 5, 1, 5, 8}
	for n := 0; n < b.N; n++ {
		MinAvgTwoSlice(A)
	}
}
