// What is the value of the first triangle number to have over 5 hundred divisors

package problem_12

import (
	"fmt"
	"math"
)

func Execute() int64 {
	//premise: if a number has
	max_factors := 500
	var factor_count int
	max_factor_count := 0
	min_triangle_num := int64(1)

	for i := int64(1); factor_count <= max_factors; i += i + 1 {
		factor_count = 0
		sqrti := int64(math.Sqrt(float64(i)))
		for j := int64(1); j <= sqrti; j++ {
			if i%j == 0 {
				factor_count += 2
				if j == sqrti {
					factor_count -= 1
				}
			}

		}
		if factor_count > max_factor_count {
			max_factor_count = factor_count
			fmt.Println(i)
			fmt.Println(factor_count)
		}
		if factor_count > max_factors {
			min_triangle_num = i
		}
	}
	return min_triangle_num
}
