// Package leap checks if the current year is a leap year
package leap

// IsLeapYear is a function that checks if the current year is a leap year
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 && year%400 != 0 {
			return false
		}
		return true
	}
	return false
}
