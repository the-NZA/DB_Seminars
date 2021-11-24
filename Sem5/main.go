package main

import "fmt"

// Input		Output
// string num char	num
// Examples:
// 1) "a" 100 'a'  -> 100
// 2) "aba" 10 'a' -> 7
// 3) "abcd" 8 'a' -> 2

func main() {
	fmt.Println(solverRight("a", 100, 'a'))
	fmt.Println(solverRight("aba", 10, 'a'))
	fmt.Println(solverRight("abcd", 8, 'a'))
	fmt.Println(solverRight("abcd", 8, 'e'))
	fmt.Println(solverRight("kmretasscityylpdhuwjirnqimlkcgxubxmsxpypgzxtenweirknjtasxtvxemtwxuarabssvqdnktqadhyktagjxoanknhgilnm", 736778906400, 'a'))
}

func solverDumb(s string, n uint, c rune) int {
	runes := []rune(s)

	// Repeat string to contain at least n chars (runes)
	for len(runes) < int(n) {
		runes = append(runes, runes...)
	}

	sum := 0

	// Count 'c' in 'runes' in first n chars
	for i := range runes[:n] {
		if runes[i] == c {
			sum++
		}
	}

	return sum
}

func solverRight(s string, n int, c rune) int {
	runes := []rune(s)
	sLen := len(runes)

	cnt := 0

	// Count 'c' in 'runes'
	for i := range runes {
		if runes[i] == c {
			cnt++
		}
	}

	fullStringCount := n / sLen * cnt

	cnt = 0
	for i := range s[:(n % sLen)] {
		if runes[i] == c {
			cnt++
		}
	}

	return fullStringCount + cnt
}
