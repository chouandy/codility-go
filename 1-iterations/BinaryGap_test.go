package iterations

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryGap(t *testing.T) {
	testCases := []struct {
		N        int
		expected int
	}{
		{
			N:        1041,
			expected: 5,
		},
		{
			N:        15,
			expected: 0,
		},
		{
			N:        32,
			expected: 0,
		},
		{
			N:        1024,
			expected: 0,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("BinaryGap[%d]", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.expected, BinaryGap(testCase.N))
		})
	}
}

func BenchmarkBinaryGap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BinaryGap(1041)
	}
}
