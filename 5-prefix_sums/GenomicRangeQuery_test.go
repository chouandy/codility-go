package prefixsums

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenomicRangeQuery(t *testing.T) {
	testCases := []struct {
		S        string
		P        []int
		Q        []int
		expected []int
	}{
		{
			S:        "CAGCCTA",
			P:        []int{2, 5, 0},
			Q:        []int{4, 5, 6},
			expected: []int{2, 4, 1},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("GenomicRangeQuery[%d]", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.expected, GenomicRangeQuery(testCase.S, testCase.P, testCase.Q))
		})
	}
}

func BenchmarkGenomicRangeQuery(b *testing.B) {
	S := "CAGCCTA"
	P := []int{2, 5, 0}
	Q := []int{4, 5, 6}
	for n := 0; n < b.N; n++ {
		GenomicRangeQuery(S, P, Q)
	}
}
