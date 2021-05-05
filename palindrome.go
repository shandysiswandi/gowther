package gowther

func IsPalindrome(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		tail := len(str) - i - 1

		if str[i] != str[tail] {
			return false
		}
	}
	return true
}
