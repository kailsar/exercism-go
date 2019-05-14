package grains

import "errors"

func Square(num int) (uint64, error) {
	var result uint64 = 1
	if num < 1 || num > 64 {
		return 0, errors.New("Square number is outside range")
	}
	return result << (uint(num) - 1), nil
}

func Total() uint64 {
	var result uint64
	return result - 1
}
