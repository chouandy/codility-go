package prefixsums

// https://app.codility.com/programmers/lessons/5-prefix_sums/min_avg_two_slice/

// MinAvgTwoSlice solution for MinAvgTwoSlice
// There is a mathematical rule that says you only have to find the min of 2 or 3 consecutive slices
func MinAvgTwoSlice(A []int) int {
	// Get A length
	n := len(A)

	// New variables
	minAvg := float64(10000)
	result := 0

	for i := 0; i < n-1; i++ {
		// The minimal of 2 consecutive slices
		if i+1 < n {
			avg := float64(A[i]+A[i+1]) / float64(2.0)
			if minAvg > avg {
				minAvg = avg
				result = i
			}
		}

		// The minimal of 3 consecutive slices
		if i+2 < n {
			avg := float64(A[i]+A[i+1]+A[i+2]) / float64(3.0)
			if minAvg > avg {
				minAvg = avg
				result = i
			}
		}
	}

	return result
}
