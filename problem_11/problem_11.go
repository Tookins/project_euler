package problem_11

import "os"

func Execute() int {

	//read data into an array of bytes 
	raw_data := make([]byte ,2000)
	raw_data, err := os.ReadFile("problem_11/data.txt")
	check(err)

	//convert raw data to a 2d array of ints
	const size = 20
	var grid [size][size]int
	i := 0
	for k := 0; k < len(raw_data); k += 3 {
		//converts two ascii chars into ints then combines them and skips the space
		//that follows
		grid[i%20][i/20] = int(raw_data[k]-48)*10+int(raw_data[k+1]-48)
		i++
	}

	// calculate the grid products and keep the largest one
	max_prod := 0
	for i:=range grid {
		for j:=range grid[0] {
			//calculate the horizontal product
			if i < size - 3 {
				prod := grid[i][j]*grid[i+1][j]*grid[i+2][j]*grid[i+3][j]
				if prod > max_prod {
					max_prod = prod
				}
			}
			//calculate the vertical product
			if j < size - 3 {
				prod := grid[i][j]*grid[i][j+1]*grid[i][j+2]*grid[i][j+3]
				if prod > max_prod {
					max_prod = prod
				}
			}
			//calculate the forward diagonal
			if i >= 3 && j < size - 3 {
				prod := grid[i][j]*grid[i-1][j+1]*grid[i-2][j+2]*grid[i-3][j+3]
				if prod > max_prod {
					max_prod = prod
				}
			}
			//calculate the back diagonal
			if i < size - 3 && j < size - 3 {
				prod := grid[i][j]*grid[i+1][j+1]*grid[i+2][j+2]*grid[i+3][j+3]
				if prod > max_prod {
					max_prod = prod
				}
			}
		}
	}

	return max_prod
}

func check(e error) {
	if e!=nil{
		panic(e)
	}
}
