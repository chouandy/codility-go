package prefixsums

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassingCars(t *testing.T) {
	testCases := []struct {
		A        []int
		expected int
	}{
		{
			A:        []int{0, 1, 0, 1, 1},
			expected: 5,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("PassingCars[%d]", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.expected, PassingCars(testCase.A))
		})
	}
}

func BenchmarkPassingCars(b *testing.B) {
	A := []int{0, 1, 0, 1, 1}
	for n := 0; n < b.N; n++ {
		PassingCars(A)
	}
}
