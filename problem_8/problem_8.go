package problem_8

import (
	"os"
)

func Execute() int64 {
	data := FileToString("problem_8/data.txt")

	var digits [1000]int
	var max_product int64 = 1

	j := 0
	for i := range data {
		if data[i] != byte('\n') && data[i] != byte('\r') {
			digits[j] = int(data[i]) - 48
			j++
		}
	}

	for i := 0; i < len(digits)-12; i++ {

		var product int64 = 1
		for j := 0; j < 13; j++ {
			product *= int64(digits[i+j])
		}
		if product > max_product {
			max_product = product
		}
	}

	return max_product
}

func FileToString(file_name string) []byte {
	raw_data, err := os.ReadFile(file_name)
	if err != nil {
		panic(err)
	}
	return raw_data
}
