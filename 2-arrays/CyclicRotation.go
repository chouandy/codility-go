package arrays

// https://app.codility.com/programmers/lessons/2-arrays/cyclic_rotation/

// CyclicRotation1 solution for CyclicRotation
func CyclicRotation1(A []int, K int) []int {
	// Get A length
	lenA := len(A)

	// Return A if
	// 1. K is 0
	// 2. A length is 0
	// 3. A length is 1
	if K == 0 || lenA == 0 || lenA == 1 {
		return A
	}

	// Update K with the remainder of the division
	if K >= lenA {
		K = K % lenA
	}

	// Return A if K is 0
	if K == 0 {
		return A
	}

	// New B by A length
	B := make([]int, lenA)

	for i, element := range A {
		// Get the moved index
		index := i + K
		// If index > A length, index minus A length
		if index >= lenA {
			index -= lenA
		}
		// Set B element
		B[index] = element
	}

	return B
}

// CyclicRotation2 solution for CyclicRotation
func CyclicRotation2(A []int, K int) []int {
	// Get A length
	lenA := len(A)

	// Return A if
	// 1. K is 0
	// 2. A length is 0
	// 3. A length is 1
	if K == 0 || lenA == 0 || lenA == 1 {
		return A
	}

	// Update K with the remainder of the division
	if K >= lenA {
		K = K % lenA
	}

	// Return A if K is 0
	if K == 0 {
		return A
	}

	// Move the right part to front
	lhs := A[lenA-K:]

	// Append the left part
	return append(lhs, A[:lenA-K]...)
}
