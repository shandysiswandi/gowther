package gowther

// Fibonacci is : https://en.wikipedia.org/wiki/Fibonacci_number
func Fibonacci(input int) []int {
	result := make([]int, 0)

	if input <= 1 {
		return append(result, 0, 1)
	}

	var temp int
	for i := 0; i < input; i++ {
		if i < 2 {
			temp = i
		} else {
			temp = result[i-1] + result[i-2]
		}

		result = append(result, temp)
	}

	return result
}
