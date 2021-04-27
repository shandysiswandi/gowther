package gowther

// FizzBuzz is : https://en.wikipedia.org/wiki/Fizz_buzz
//

func FizzBuzz(input int) []interface{} {
	result := make([]interface{}, 0)

	for i := 1; i <= input; i++ {
		if IsFizzBuzz(i) {
			result = append(result, "Fizz Buzz")
		} else if IsFizz(i) {
			result = append(result, "Fizz")
		} else if IsBuzz(i) {
			result = append(result, "Buzz")
		} else {
			result = append(result, i)
		}
	}

	return result
}

func IsFizz(input int) bool {
	return input%3 == 0
}

func IsBuzz(input int) bool {
	return input%5 == 0
}

func IsFizzBuzz(input int) bool {
	return input%5 == 0 && input%3 == 0
}
