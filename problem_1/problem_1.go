package problem_1

// return the sum of all the multiples of 3 or 5 
func Problem01 (ceil int) int {
	sum := 0
	i := 0
	for i < ceil {
		if (i % 3 == 0 || i % 5 == 0) {
			sum += i
		}
		i++
	}
	return sum
}
