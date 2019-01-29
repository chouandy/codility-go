package arrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCyclicRotation(t *testing.T) {
	testCases := []struct {
		A        []int
		K        int
		expected []int
	}{
		{
			A:        []int{3, 8, 9, 7, 6},
			K:        3,
			expected: []int{9, 7, 6, 3, 8},
		},
		{
			A:        []int{0, 0, 0},
			K:        1,
			expected: []int{0, 0, 0},
		},
		{
			A:        []int{1, 2, 3, 4},
			K:        4,
			expected: []int{1, 2, 3, 4},
		},
		{
			A:        []int{},
			K:        0,
			expected: []int{},
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("CyclicRotation[%d]", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.expected, CyclicRotation1(testCase.A, testCase.K))
			assert.Equal(t, testCase.expected, CyclicRotation2(testCase.A, testCase.K))
		})
	}
}

func BenchmarkCyclicRotation1(b *testing.B) {
	A := []int{3, 8, 9, 7, 6}
	K := 3
	for n := 0; n < b.N; n++ {
		CyclicRotation1(A, K)
	}
}

func BenchmarkCyclicRotation2(b *testing.B) {
	A := []int{3, 8, 9, 7, 6}
	K := 3
	for n := 0; n < b.N; n++ {
		CyclicRotation2(A, K)
	}
}
