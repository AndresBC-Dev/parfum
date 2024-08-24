package math

import "strconv"

func Sum(a, b int) int {
	return a + b
}

func sumString(a, b string) int {
	c, err := strconv.Atoi(a)
	if err != nil {
		return 0
	}

	d, err := strconv.Atoi(b)
	if err != nil {
		return 0
	}

	return c + d
}
