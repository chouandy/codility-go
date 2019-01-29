package countingelements

// https://app.codility.com/programmers/lessons/4-counting_elements/missing_integer/

// MissingInteger solution for MissingInteger
func MissingInteger(A []int) int {
	// New variables
	seen := make(map[int]bool)
	max := 0

	for _, element := range A {
		// Check duplicate element
		if seen[element] {
			continue
		}
		// Set seen
		seen[element] = true
		// Update max
		if max < element {
			max = element
		}
	}

	// Find missing integer with max + 1
	for i := 1; i <= max+1; i++ {
		if !seen[i] {
			return i
		}
	}

	return 1
}
