package prefixsums

// https://app.codility.com/programmers/lessons/5-prefix_sums/passing_cars/

// PassingCars solution for PassingCars
func PassingCars(A []int) int {
	// Get prefix sums
	sums := passingCarsPrefixSums(A)

	// New variables
	count := 0

	for i, element := range A {
		// Count passing if element = 0
		if element == 0 {
			count += passingCarsCountTotal(sums, i, len(A)-1)
		}
	}

	// Return -1 if pairs of passing cars exceeds 1,000,000,000
	if count > 1000000000 {
		return -1
	}

	return count
}

func passingCarsPrefixSums(A []int) []int {
	n := len(A)
	sums := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		sums[i] = sums[i-1] + A[i-1]
	}
	return sums
}

func passingCarsCountTotal(sums []int, x, y int) int {
	return sums[y+1] - sums[x]
}
