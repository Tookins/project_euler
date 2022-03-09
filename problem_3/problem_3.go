package problem_3

import (
	"math"
	"os"
	"strconv"
)

const file_name string = "problem_3/data.txt"

func Execute() int64 {

	in, err := os.ReadFile(file_name)

	check(err)

	data, err := strconv.ParseInt(string(in), 10, 64)

	check(err)

	var greatest_prime_factor int64 = 0

	for i := int64(2); i <= int64(math.Sqrt(float64(data))); i++ {

		remainder := i
		dividend := data / i

		if data%i == 0 {
			if dividend > greatest_prime_factor && IsPrime(dividend) {
				greatest_prime_factor = dividend
			}
			if remainder > greatest_prime_factor && IsPrime(remainder) {
				greatest_prime_factor = remainder
			}

		}

	}

	return greatest_prime_factor
}

func IsPrime(num int64) bool {

	var i int64
	for i = 2; i <= int64(math.Sqrt(float64(num))); i++ {

		if num%i == 0 {

			return false

		}
	}

	return true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
