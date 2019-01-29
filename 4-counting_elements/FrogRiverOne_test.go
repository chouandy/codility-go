package countingelements

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrogRiverOne(t *testing.T) {
	testCases := []struct {
		X        int
		A        []int
		expected int
	}{
		{
			X:        5,
			A:        []int{1, 3, 1, 4, 2, 3, 5, 4},
			expected: 6,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("FrogRiverOne[%d]", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.expected, FrogRiverOne1(testCase.X, testCase.A))
			assert.Equal(t, testCase.expected, FrogRiverOne2(testCase.X, testCase.A))
		})
	}
}

func BenchmarkFrogRiverOne1(b *testing.B) {
	X := 5
	A := []int{1, 3, 1, 4, 2, 3, 5, 4}
	for n := 0; n < b.N; n++ {
		FrogRiverOne1(X, A)
	}
}

func BenchmarkFrogRiverOne2(b *testing.B) {
	X := 5
	A := []int{1, 3, 1, 4, 2, 3, 5, 4}
	for n := 0; n < b.N; n++ {
		FrogRiverOne2(X, A)
	}
}
