package countingelements

// https://app.codility.com/programmers/lessons/4-counting_elements/frog_river_one/

// FrogRiverOne1 solution for FrogRiverOne
func FrogRiverOne1(X int, A []int) int {
	// New variables
	seen := make(map[int]bool)
	count := 0

	for i, element := range A {
		// Check duplicate element
		if seen[element] {
			continue
		}
		// Set seen
		seen[element] = true
		// Increase count
		count++
		// Check count and X
		if count == X {
			return i
		}
	}

	return -1
}

// FrogRiverOne2 solution for FrogRiverOne
func FrogRiverOne2(X int, A []int) int {
	// New variables
	seen := make(map[int]bool)
	sumA := 0

	// Sum X
	sumX := X * (X + 1) / 2

	for i, element := range A {
		// Check duplicate element
		if seen[element] {
			continue
		}
		// Set seen
		seen[element] = true
		// Add sum A and element up
		sumA += element
		// Check sum A and sum X
		if sumA == sumX {
			return i
		}
	}

	return -1
}
