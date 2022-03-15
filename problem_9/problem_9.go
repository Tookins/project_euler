package problem_9

func Execute() int {
	// sum of desired pythagorean triplet (a, b, c) st a+b+c=data
	data := 1000
	// product of a b and c
	var result int
	//
	for a := 2; a < data; a++ {
		for b := 1; b < a; b++ {
			c := 1000 - a - b
			if a*a+b*b == c*c {
				result = a * b * c
			}
		}
	}
	return result
}
