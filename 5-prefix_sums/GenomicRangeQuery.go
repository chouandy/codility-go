package prefixsums

// https://app.codility.com/programmers/lessons/5-prefix_sums/genomic_range_query/

// GenomicRangeQuery solution for GenomicRangeQuery
func GenomicRangeQuery(S string, P []int, Q []int) []int {
	// Get prefix sums
	sumsA, sumsC, sumsG := genomicRangeQueryPrefixSums(S)

	// New result
	result := make([]int, len(P))

	for i := range P {
		if genomicRangeQueryCountTotal(sumsA, P[i], Q[i]) > 0 {
			result[i] = 1
		} else if genomicRangeQueryCountTotal(sumsC, P[i], Q[i]) > 0 {
			result[i] = 2
		} else if genomicRangeQueryCountTotal(sumsG, P[i], Q[i]) > 0 {
			result[i] = 3
		} else {
			result[i] = 4
		}
	}

	return result
}

func genomicRangeQueryPrefixSums(S string) ([]int, []int, []int) {
	// Get S length
	n := len(S)
	// New A, C, G prefix sums
	sumsA := make([]int, n+1)
	sumsC := make([]int, n+1)
	sumsG := make([]int, n+1)

	// Get A, C, G prefix sums of the number of occurences of each letter from S
	for i := 1; i < n+1; i++ {
		sumsA[i] = sumsA[i-1]
		sumsC[i] = sumsC[i-1]
		sumsG[i] = sumsG[i-1]

		letter := string(S[i-1])
		if letter == "A" {
			sumsA[i]++
		} else if letter == "C" {
			sumsC[i]++
		} else if letter == "G" {
			sumsG[i]++
		}
	}

	return sumsA, sumsC, sumsG
}

func genomicRangeQueryCountTotal(sums []int, x, y int) int {
	return sums[y+1] - sums[x]
}
