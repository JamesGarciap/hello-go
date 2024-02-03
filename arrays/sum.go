package arrays

func Sum(numbers []int) (sum int) {
	sum = 0

	for _, number := range numbers {
		sum += number
	}

	return
}

func SumAll(numbersToSum ...[]int) []int {
	lengtOfNumbers := len(numbersToSum)
	sums := make([]int, lengtOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
