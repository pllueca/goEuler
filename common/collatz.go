package common

// Returns the number of collatz steps
// Slow version
func CollatzSteps(n int) int {
	var s int
	for s=0; n != 1;{
		if s%2 == 0 {
			n /= 2
		} else {
			n = 3 * n + 1
		}
		s += 1
	}
	return s
}

// Returns the number of collatz steps
// Fast version, uses a lookuptable to store results
// initially lookup = map[1:1]
func DynamicCollatzSteps(n int, lookup map[int]int) int {
	s, exists := lookup[n]
	if exists {
		return s
	}

	var v int
	if s%2 != 0 {
		v = 1 + DynamicCollatzSteps((3*n + 1), lookup)
	} else {
		v = 1 + DynamicCollatzSteps(n/2, lookup)
	}
	
	lookup[n] = v
	return v
}