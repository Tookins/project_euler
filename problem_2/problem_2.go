package problem_2

import (
	"os"
	"strconv"
)

const file_name = "problem_2/data.txt"

func Execute() int64 {

	in, err := os.ReadFile(file_name)

	check(err)

	data, err := strconv.Atoi(string(in))

	check(err)

	ceil := int64(data)

	fib_0 := int64(0)

	fib_1 := int64(1)

	result := int64(0)

	for fib_1 < ceil {

		fib_0 += fib_1

		fib_0, fib_1 = fib_1, fib_0

		if fib_1%2 == 0 {

			result += fib_1

		}

	}

	return result
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
