// Package raindrops checks numbers for their factors and depending
// on the results, outputs different words
package raindrops

import "strconv"

// Convert is a function that converts numbers in to the words "Pling", "Plang"
// and "Plong" depending on the number's factors.
func Convert(number int) string {
	var result string
	if number%3 == 0 {
		result += "Pling"
	}
	if number%5 == 0 {
		result += "Plang"
	}
	if number%7 == 0 {
		result += "Plong"
	}
	if result == "" {
		result = strconv.Itoa(number)
	}
	return result
}
