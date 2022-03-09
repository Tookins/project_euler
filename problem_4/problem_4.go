package problem_4

const file_name string = "problem_3/data.txt"

func Execute() int {
	/*
		in, err := os.ReadFile(file_name)

		check(err)

		data, err := strconv.Atoi(string(in))

		check(err)
	*/

	var product int
	var max_palindrome int

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			product = i * j
			if product > max_palindrome && isPalindrome(product) {
				max_palindrome = product
			}
		}
	}
	return max_palindrome
}

func isPalindrome(number int) bool {
	var digits []int
	for temp := number; temp > 0; temp /= 10 {
		digits = append(digits, temp%10)
	}
	for i := 0; i <= len(digits)/2; i++ {
		if digits[i] != digits[len(digits)-i-1] {
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
