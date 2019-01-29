package arrays

// https://app.codility.com/programmers/lessons/3-time_complexity/perm_missing_elem/

// PermMissingElem solution for PermMissingElem
func PermMissingElem(A []int) int {
	// Get real array length
	lenReal := len(A) + 1
	// Get real array sum
	sumReal := (1 + lenReal) * lenReal / 2
	// Sum A
	sumA := 0
	for _, element := range A {
		sumA += element
	}

	// Return real array sum minus sum A
	return sumReal - sumA
}
