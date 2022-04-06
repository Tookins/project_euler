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
	min_triangle_num := int64(1)

	for i:=int64(1); factor_count <= max_factors; i++ {
		factor_count = 0
		for j:=int64(1); float64(j) <= math.Sqrt(float64(i)); j++ {
			if i%j == 0 {
				factor_count += 2
			}
			//debugging
			if factor_count > 100 {
				fmt.Println(i)
			}
		}
		if factor_count > max_factors {
			min_triangle_num = i
		}
	}
	return min_triangle_num
}
