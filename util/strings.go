package util

import "strconv"

func Levenshtein(s1, s2 string) int {
	diff := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff++
		}
	}
	return diff
}

func AsciiValue(s string) int {
	return int(s[0])
}

func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
