package main

import(
	"fmt"
	"os"
	"strconv"
//	"project_euler/problem_1"
)

//error checker for files
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// generalize the process of running the package containing the algorithm
// for a given problem and providing the correct data to the algorithm

func main() {

	// parse clarg with format: euler [problem number]
	problem := os.Args[1]
	
	switch problem {
	
		case "1": {
			fstring := "/home/usetheforc3luke/go/src/project_euler/problem_1/data.txt"
			in, err := os.ReadFile(fstring)

			check(err)

			result, err := strconv.Atoi(string(in[:len(in) - 1]))

			check(err)

			fmt.Println(result)

			//result := problem_1.Problem01(n)
		}
		default: fmt.Println("Problem number not found.")
	}
}
