package main

import (
	"fmt"
	"strings"
)

func Valid(cardNumber string) bool {
	total := 0
	cardNumber = strings.ReplaceAll(cardNumber, " ", "")
	length := len(cardNumber)
	if length < 2 {
		return false
	}
	for x := 0; x < length; x++ {
		fmt.Println(cardNumber[length-(x+1)])
		intDigit := int(cardNumber[length-(x+1)] - 48)
		fmt.Println(intDigit)
		if intDigit < 0 || intDigit > 94 {

			return false
		}

		if x%2 != 0 {
			intDigit = intDigit * 2
			if intDigit > 9 {
				intDigit = intDigit - 9
			}
		}
		fmt.Println(intDigit)
		total += intDigit
	}
	if total%10 == 0 {
		return true
	}
	return false
}
func main() {

	Valid("059")
}
