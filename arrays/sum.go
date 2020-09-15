package main

func Sum(numbers []int) (sum int)  {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(arrays ...[]int) (sumArray []int) {
	for _, array := range arrays {
		sumArray = append(sumArray, Sum(array))
	}
	return
}

func SumAllTails(arrays ...[]int) (sumArray []int) {
	for _, array := range arrays {
		tail := array[1:]
		sumArray = append(sumArray, Sum(tail))
	}
	return
}
