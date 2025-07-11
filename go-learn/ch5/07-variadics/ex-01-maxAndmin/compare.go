package main

import "fmt"

func Max(v ...int) (int, error) {
	if len(v) == 0 {
		return 0, fmt.Errorf("max: no argument")
	}
	max := v[0]
	for _, val := range v {
		if val > max {
			max = val
		}
	}
	return max, nil
}

func Min(v ...int) (int, error) {
	if len(v) == 0 {
		return 0, fmt.Errorf("min: no argument")
	}
	min := v[0]
	for _, val := range v {
		if val < min {
			min = val
		}
	}
	return min, nil
}
