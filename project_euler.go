package main

import (
	"fmt"
	"os"

	"project_euler/problem_1"
)

// generalize the process of running the package containing the algorithm
// for a given problem and providing the correct data to the algorithm

func main() {

	// parse clarg with format: euler [problem number]
	problem := os.Args[1]

	switch problem {

	case "1":
		fmt.Println(problem_1.Execute())
	default:
		fmt.Println("Problem number not found.")
	}
}
