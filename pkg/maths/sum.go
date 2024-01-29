package maths

import "time"

func Sum[T int | time.Duration | float64 | uint](data []T) T {
	var result T
	for _, v := range data {
		result = result + v
	}
	return result
}
