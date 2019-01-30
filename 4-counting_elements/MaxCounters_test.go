package countingelements

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxCounters(t *testing.T) {
	testCases := []struct {
		N        int
		A        []int
		expected []int
	}{
		{
			N:        5,
			A:        []int{3, 4, 4, 6, 1, 4, 4},
			expected: []int{3, 2, 2, 4, 2},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("MaxCounters[%d]", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.expected, MaxCounters(testCase.N, testCase.A))
		})
	}
}

func BenchmarkMaxCounters(b *testing.B) {
	N := 5
	A := []int{3, 4, 4, 6, 1, 4, 4}
	for n := 0; n < b.N; n++ {
		MaxCounters(N, A)
	}
}
