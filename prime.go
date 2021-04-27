package gowther

// https://en.wikipedia.org/wiki/Prime_number
//

func Primes(input int) []int {
	result := make([]int, 0)
	for i := 1; i <= input; i++ {
		if IsPrime(i) {
			result = append(result, i)
		}
	}

	return result
}

func IsPrime(input int) bool {
	for i := 2; i < input; i++ {
		if input%i == 0 {
			return false
		}
	}

	return input > 1
}
