package countingelements

// https://app.codility.com/programmers/lessons/4-counting_elements/perm_check/

// PermCheck solution for PermCheck
func PermCheck(A []int) int {
	// New variables
	seen := make(map[int]bool)

	// Sum A
	sumA := 0
	for _, element := range A {
		// Check duplicate element
		if seen[element] {
			return 0
		}
		// Set seen
		seen[element] = true
		// Add sum A and element up
		sumA += element
	}

	// Get A length
	lenA := len(A)

	// Check sum A
	if (1+lenA)*lenA/2 == sumA {
		return 1
	}

	return 0
}
