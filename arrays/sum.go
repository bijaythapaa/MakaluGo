package arrays

func Sum(numbers [5]int) int {
	sum := 0
	//for i := 0; i < 5; i++ {
	//	sum += numbers[i]
	//}

	for _, num := range numbers {
		sum += num
	}
	return sum
}
