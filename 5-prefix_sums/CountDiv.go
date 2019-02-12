package prefixsums

import "math"

// https://app.codility.com/programmers/lessons/5-prefix_sums/count_div/

// CountDiv solution for CountDiv
func CountDiv(A int, B int, K int) int {
	// The quotient of B/K - the quotient of A/K
	divisibles := int(math.Floor(float64(B/K)) - math.Floor(float64(A/K)))

	// Check A is divisible by K or not
	if A%K == 0 {
		divisibles++
	}

	return divisibles
}
