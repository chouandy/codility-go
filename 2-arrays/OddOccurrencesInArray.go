package arrays

// https://app.codility.com/programmers/lessons/2-arrays/odd_occurrences_in_array/

// OddOccurrencesInArray solution for OddOccurrencesInArray
func OddOccurrencesInArray(A []int) int {
	// New variables
	var histogram = make(map[int]int)

	// Group by element
	for _, element := range A {
		histogram[element]++
	}

	// Find unpaired element
	for k, v := range histogram {
		if v%2 == 1 {
			return k
		}
	}

	return -1
}
