// TODO
// add in flags for command line arguments to do the following
// print a problem statement to the console
// read input from stdin
// measure performance of a solution

package main

import (
	"fmt"
	"os"
	"strconv"

	"project_euler/problem_1"
	"project_euler/problem_2"
	"project_euler/problem_3"
	"project_euler/problem_4"

	"github.com/dmars8047/eulertools"
)

// generalize the process of running the package containing the algorithm
// for a given problem and providing the correct data to the algorithm

const user_agent string = "RJO6969"

func main() {

	if len(os.Args) < 2 {
		fmt.Printf("\n\nError: Please enter the problem number to be executed.\n\n")
		return
	}

	problem, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Printf("\n\nSupplied problem number: %s is not a number\n\n", os.Args[1])
		return
	}

	eulertools.PrintPremise(problem, user_agent)

	switch problem {

	case 1:
		fmt.Println(problem_1.Execute())

	case 2:
		fmt.Println(problem_2.Execute())

	case 3:
		fmt.Println(problem_3.Execute())

	case 4:
		fmt.Println(problem_4.Execute())

	default:
		fmt.Println("Problem number not found.")
	}
}
