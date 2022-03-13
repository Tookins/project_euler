package problem_5

import (
	"fmt"
	"math"
	"os"
	"project_euler/problem_3"
	"strconv"
)

func main() {
	Execute()
}

func Execute() int64 {

	//read an integer from data.txt
	file_name := "problem_5/data.txt"
	data := FileToInt(file_name)
	fmt.Println(data)

	//make a map whose keys are the primes less than the upper bound
	primes := make(map[int]int)
	for i := 2; i <= data; i++ {
		if problem_3.IsPrime(int64(i)) {
			primes[i] = 0
		}
	}

	//iterate over the primes and raise them to the nth power
	for key, _ := range primes {
		temp := key
		for temp <= data {
			temp *= key
			primes[key] += 1
		}
	}

	//calculate the product of each prime factor raised to its multiplicity
	var product int64 = 1
	for key, val := range primes {
		product *= int64(math.Pow(float64(key), float64(val)))
	}

	return product
}

func FileToInt(file string) int {

	raw_data, err := os.ReadFile(file)

	Check(err)

	data, err := strconv.Atoi(string(raw_data))

	Check(err)

	return data
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
