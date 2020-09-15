package main

func Sum(numbers []int) (sum int)  {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(arrays ...[]int) (sumArray []int) {
	var sum int
	for _, array := range arrays {
		sum = 0
		for _, n := range array {
			sum += n
		}
		sumArray = append(sumArray, sum)
	}
	return
}
