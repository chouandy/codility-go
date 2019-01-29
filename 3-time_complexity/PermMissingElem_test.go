package arrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermMissingElem(t *testing.T) {
	testCases := []struct {
		A        []int
		expected int
	}{
		{
			A:        []int{2, 3, 1, 5},
			expected: 4,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("PermMissingElem[%d]", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.expected, PermMissingElem(testCase.A))
		})
	}
}

func BenchmarkPermMissingElem(b *testing.B) {
	A := []int{2, 3, 1, 5}
	for n := 0; n < b.N; n++ {
		PermMissingElem(A)
	}
}
