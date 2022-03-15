package problem_10

func Execute() int {
	//values will contain first 2000000 nat numbers
	//is_prime[i] will be the result of IsPrime(i)
	var values [2000001]int
	var not_prime [2000001]bool
	for i := range values {
		values[i] = i
	}
	// loop over the values and set values of not_prime
	// accordingly (seive of aristophenes??)
	not_prime[0] = true
	not_prime[1] = true //array includes 0 and 1 which are not prime by convention
	for i := 2; i < len(values)-1; i++ {
		for j := i * i; j < len(values); j += i {
			not_prime[j] = true
		}
	}
	// calculate the sum of the primes
	sum := 0
	for i, j := range not_prime {
		if !j {
			sum += values[i]
		}
	}
	return sum
}
