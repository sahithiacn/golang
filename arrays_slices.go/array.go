package main

func Sum(numbers []int) int {
	//return 0
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum

}
