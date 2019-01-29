package timecomplexity

import "math"

// https://app.codility.com/programmers/lessons/3-time_complexity/tape_equilibrium/

// TapeEquilibrium solution for TapeEquilibrium
func TapeEquilibrium(A []int) int {
	// Get first sum left and sum right
	sumLeft := A[0]
	sumRight := 0
	for _, element := range A[1:] {
		sumRight += element
	}

	// Set min with first diff
	min := int(math.Abs(float64(sumLeft - sumRight)))

	// Start from 1 index element, and end with N - 1 element
	for i := 1; i < len(A)-1; i++ {
		// Add sum left and element up
		sumLeft += A[i]
		// Subtract element from sum right
		sumRight -= A[i]
		// Get diff
		diff := int(math.Abs(float64(sumLeft - sumRight)))
		// Update min if diff < min
		if diff < min {
			min = diff
		}
	}

	return min
}
