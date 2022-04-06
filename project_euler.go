// TODO

// problem 5 broken somewhere maybe ISPRIME

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
	"project_euler/problem_10"
	"project_euler/problem_2"
	"project_euler/problem_3"
	"project_euler/problem_4"
	"project_euler/problem_5"
	"project_euler/problem_6"
	"project_euler/problem_7"
	"project_euler/problem_8"
	"project_euler/problem_9"
	"project_euler/problem_11"
	"project_euler/problem_12"
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

	//eulertools.PrintPremise(problem, user_agent)

	switch problem {

	case 1:
		fmt.Println(problem_1.Execute())

	case 2:
		fmt.Println(problem_2.Execute())

	case 3:
		fmt.Println(problem_3.Execute())

	case 4:
		fmt.Println(problem_4.Execute())

	case 5:
		fmt.Println(problem_5.Execute())

	case 6:
		fmt.Println(problem_6.Execute())

	case 7:
		fmt.Println(problem_7.Execute())

	case 8:
		fmt.Println(problem_8.Execute())

	case 9:
		fmt.Println(problem_9.Execute())

	case 10:
		fmt.Println(problem_10.Execute())

	case 11:
		fmt.Println(problem_11.Execute())

	case 12:
		fmt.Println(problem_12.Execute())

	default:
		fmt.Println("Problem number not found.")
	}
}
