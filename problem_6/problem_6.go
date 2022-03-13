package problem_6

import (
	"project_euler/problem_5"
)

func Execute() int64 {

	var data int64 = int64(problem_5.FileToInt("problem_6/data.txt"))

	sum := int64(0)
	sum_of_squares := int64(0)

	for i := int64(1); i <= data; i++ {
		sum_of_squares += i * i
		sum += i
	}

	return sum*sum - sum_of_squares
}
