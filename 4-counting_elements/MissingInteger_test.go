package countingelements

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingInteger(t *testing.T) {
	testCases := []struct {
		A        []int
		expected int
	}{
		{
			A:        []int{1, 3, 6, 4, 1, 2},
			expected: 5,
		},
		{
			A:        []int{1, 2, 3},
			expected: 4,
		},
		{
			A:        []int{-1, -3},
			expected: 1,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("MissingInteger[%d]", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.expected, MissingInteger(testCase.A))
		})
	}
}

func BenchmarkMissingInteger(b *testing.B) {
	A := []int{1, 3, 6, 4, 1, 2}
	for n := 0; n < b.N; n++ {
		MissingInteger(A)
	}
}
