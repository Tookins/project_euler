package problem_5

cont file_name string = "problem_5/data.txt"

func Execute () int {
	data := FileToInt(file_name)
	

	
	for i:= 2; i <= data; i++ {

}

func (div int) []int {
	
}

func FileToInt (file string) int {

	raw_data, err = os.ReadFile(file_name)
	
	check(err)

	data, err = strconv.Atoi(raw_data)

	check(err)

}

func  check (e error) {
	if e != nil {
		panic(e)
	}
}
