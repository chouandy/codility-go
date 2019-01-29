package timecomplexity

import "math"

// https://app.codility.com/programmers/lessons/3-time_complexity/frog_jmp/

// FrogJmp1 solution for FrogJmp
func FrogJmp1(X int, Y int, D int) int {
	// Get the distance between X and Y
	distance := Y - X

	// Get the jumps with D
	jumps := distance / D

	// Return jumps if the remainder of distance divided by D is 0
	if distance%D == 0 {
		return jumps
	}

	// Return jumps + 1 if the remainder of distance divided by D is not 0
	return jumps + 1
}

// FrogJmp2 solution for FrogJmp
func FrogJmp2(X int, Y int, D int) int {
	return int(math.Ceil(float64(float64(Y-X) / float64(D))))
}
