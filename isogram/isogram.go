package isogram

import "strings"

func IsIsogram(word string) bool {
	m := map[string]bool{"": false}
	for _, runeValue := range word {
		s1 := strings.ToLower(string(runeValue))
		if m[s1] == true {
			return false
		}
		if s1 == " " || s1 == "-" {
			continue
		}
		m[s1] = true
	}
	return true
}
