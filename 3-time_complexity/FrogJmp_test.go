package arrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrogJmp(t *testing.T) {
	testCases := []struct {
		X        int
		Y        int
		D        int
		expected int
	}{
		{
			X:        10,
			Y:        85,
			D:        30,
			expected: 3,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("FrogJmp[%d]", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.expected, FrogJmp1(testCase.X, testCase.Y, testCase.D))
			assert.Equal(t, testCase.expected, FrogJmp2(testCase.X, testCase.Y, testCase.D))
		})
	}
}

func BenchmarkFrogJmp1(b *testing.B) {
	X := 10
	Y := 85
	D := 30
	for n := 0; n < b.N; n++ {
		FrogJmp1(X, Y, D)
	}
}

func BenchmarkFrogJmp2(b *testing.B) {
	X := 10
	Y := 85
	D := 30
	for n := 0; n < b.N; n++ {
		FrogJmp2(X, Y, D)
	}
}
