package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 34.5, 39.9, 40.1}

	groups := groupTemperatures(temperatures)

	for key, values := range groups {
		fmt.Printf("%d: %v\n", key, values)
	}
}

func groupTemperatures(temps []float64) map[int][]float64 {
	groups := make(map[int][]float64)

	for _, temp := range temps {
		groupKey := getGroupKey(temp)

		groups[groupKey] = append(groups[groupKey], temp)
	}

	return groups
}

func getGroupKey(temp float64) int {
	if temp >= 0 {
		return int(math.Floor(temp/10)) * 10
	}
	return int(math.Ceil(temp/10)) * 10
}
