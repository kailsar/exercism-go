// Package luhn checks the validity of a string according to the Luhn formula
package luhn

import (
	"strings"
)

// Valid checks the validity of a cardNumber string according to the Luhn formula
func Valid(cardNumber string) bool {
	total := 0
	cardNumber = strings.ReplaceAll(cardNumber, " ", "")
	length := len(cardNumber)
	if length < 2 {
		return false
	}
	for x := 0; x < length; x++ {
		intDigit := int(cardNumber[length-(x+1)] - 48) // Get the int digit from ASCII encoded byte
		if intDigit < 0 || intDigit > 9 {
			return false
		}

		if x%2 != 0 {
			intDigit = intDigit * 2
			if intDigit > 9 {
				intDigit = intDigit - 9
			}
		}
		total += intDigit
	}
	if total%10 == 0 {
		return true
	}
	return false
}
