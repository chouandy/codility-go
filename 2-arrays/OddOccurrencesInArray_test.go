package arrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOddOccurrencesInArray(t *testing.T) {
	testCases := []struct {
		A        []int
		expected int
	}{
		{
			A:        []int{9, 3, 9, 3, 9, 7, 9},
			expected: 7,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("OddOccurrencesInArray[%d]", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.expected, OddOccurrencesInArray(testCase.A))
		})
	}
}
