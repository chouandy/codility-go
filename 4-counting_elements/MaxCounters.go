package countingelements

// https://app.codility.com/programmers/lessons/4-counting_elements/max_counters/

// MaxCounters solution for MaxCounters
func MaxCounters(N int, A []int) []int {
	// New variables
	counters := make([]int, N)
	maxNum := N + 1
	base := 0
	currentMax := 0

	for _, element := range A {
		// Check element is max number or not
		if element == maxNum {
			// Update base number to current max number
			base = currentMax
		} else {
			// Get counter index
			index := element - 1
			// Update counter to base if counter number < base
			if counters[index] < base {
				counters[index] = base
			}
			// Increase counter
			counters[index]++
			// Update current max number
			if currentMax < counters[index] {
				currentMax = counters[index]
			}
		}
	}

	// Update counters to base, if counter number < base number
	for i, v := range counters {
		if v < base {
			counters[i] = base
		}
	}

	return counters
}
