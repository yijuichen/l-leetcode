package main

import "fmt"

/*
思路：
	1. 被限定o(1)
	2. 頭尾交換

*/

func reverseString(s []byte) {
	i := 0
	j := len(s) - 1

	for i < j {
		s[i], s[j] = s[j], s[i]
		j--
		i++
	}
	fmt.Print(string(s))

}

func main() {
	reverseString([]byte{'h', 'e', 'l', 'l', 'o'})
}
