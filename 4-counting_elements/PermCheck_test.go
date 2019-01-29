package countingelements

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermCheck(t *testing.T) {
	testCases := []struct {
		A        []int
		expected int
	}{
		{
			A:        []int{4, 1, 3, 2},
			expected: 1,
		},
		{
			A:        []int{4, 1, 3},
			expected: 0,
		},
		{
			A:        []int{9, 5, 7, 3, 2, 7, 3, 1, 10, 8},
			expected: 0,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("PermCheck[%d]", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.expected, PermCheck(testCase.A))
		})
	}
}

func BenchmarkPermCheck(b *testing.B) {
	A := []int{4, 1, 3, 2}
	for n := 0; n < b.N; n++ {
		PermCheck(A)
	}
}
