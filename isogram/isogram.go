// Package isogram checks if a given word is an isogram, i.e. contains no
// repeat letters
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram checks if a string contains the same alphabetic
// character(ignoring case) more than once.
func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	used := make(map[rune]bool)
	for _, letter := range word {
		if unicode.IsLetter(letter) {
			if used[letter] {
				return false
			}
			{
				used[letter] = true
			}
		}

	}
	return true
}
