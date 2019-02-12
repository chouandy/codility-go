package prefixsums

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountDiv(t *testing.T) {
	testCases := []struct {
		A        int
		B        int
		K        int
		expected int
	}{
		{
			A:        6,
			B:        11,
			K:        2,
			expected: 3,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("CountDiv[%d]", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.expected, CountDiv(testCase.A, testCase.B, testCase.K))
		})
	}
}

func BenchmarkCountDiv(b *testing.B) {
	A := 6
	B := 11
	K := 2
	for n := 0; n < b.N; n++ {
		CountDiv(A, B, K)
	}
}
