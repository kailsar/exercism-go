// Package grains calculates how many grains are on a given chess square, when the
// number starts at 1, and doubles for each square, and how many are on the board
// in total

package grains

import "errors"

// Square calculates the number of grains on a given square
func Square(num int) (uint64, error) {
	var result uint64 = 1
	if num < 1 || num > 64 {
		return 0, errors.New("Square number is outside range")
	}
	for x := 1; x < num; x++ {
		result = result * 2
	}
	return result, nil
}

// Total calculates the total number of grains on the board
func Total() uint64 {
	var result uint64
	for x := 1; x <= 64; x++ {
		square, _ := Square(x)
		result = result + square
	}
	return result
}
