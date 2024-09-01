package main

func validPalindrome(s string) bool {
	i, j := 0, len(s)-1

	for i < j {
		if s[i] != s[j] {
			// 跳過左邊或跳過右邊
			return isPalindrome(s, i+1, j) || isPalindrome(s, i, j-1)
		}
		i++
		j--
	}

	return true
}

func isPalindrome(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
