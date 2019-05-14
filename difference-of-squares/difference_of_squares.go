// Package isogram calculates the difference between the square of the sum of the
// integers from 1 to n, and the sum of the squares of the integers from 1 to n

package diffsquares

// SquareOfSum calculates the square of the sum of the integers from 1 to n
func SquareOfSum(n int) int {
	total := (n * (n + 1)) / 2
	return total * total
}

// SumOfSquares calculates the sum of the squares of the integers from 1 to n
func SumOfSquares(n int) int {
	return (n * (n + 1) * (n*2 + 1)) / 6
}

// Difference calculates the difference between the square of sum and the sum of
// squares
func Difference(number int) int {
	return SquareOfSum(number) - SumOfSquares(number)
}
