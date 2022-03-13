package problem_7

import "project_euler/problem_3"

func Execute() int {
	n := 10001
	ith_prime := 2
	for i := 1; i < n; {
		ith_prime += 1
		if problem_3.IsPrime(int64(ith_prime)) {
			i++
		}
	}
	return ith_prime
}
