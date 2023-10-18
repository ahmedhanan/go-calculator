package calculator

import "errors"

func Add(a, b float64) float64 {
	return a + b
}

func Subsctract(a, b float64) float64 {
	return a - b

}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("cannot divide 0")
	}
	return a / b, nil
}
