package integers

import (
	"errors"
	"math"
)

func Add(x int, y int) int {
	return x + y
}

func Subtracts(x, y int) int {
	return x - y
}
func Multiplies(x, y int) int {
	return x * y
}

func Divides(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return x / y, nil
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Math: negative number passed to Sqrt")
	}
	return math.Sqrt(x), nil
}
