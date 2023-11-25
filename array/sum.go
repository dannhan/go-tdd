package main

// Takes a slice of integers and returns the sum of all the elements in that slice
func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return
}

// Takes multiple slices of integers and returns a new slice containing
// the sums of the elements in each of the provided slices.
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// Takes multiple slices of integers and returns a new slice containing
// the sums of the tail elements in each of the provided slices.
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
