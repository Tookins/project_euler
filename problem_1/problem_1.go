package problem_1

import (
	"os"
	"strconv"
)

const fstring = "./data.txt"

// return the sum of all the multiples of 3 or 5
func Execute() int {

	in, err := os.ReadFile(fstring)

	check(err)

	data, err := strconv.Atoi(string(in[:len(in)-1]))

	check(err)

	sum := 0

	i := 0

	for i < data {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
		i++
	}
	return sum
}

//error checker for files
func check(e error) {
	if e != nil {
		panic(e)
	}
}
