package main

import "fmt"

type cost struct {
	day   int
	value float64
}

func getDailyCost(costs []cost) []float64 {
	costByDay := []float64{}
	for _, cost := range costs {
		if cost.day > len(costByDay)-1 {
			costByDay = append(costByDay, 0.00)
		}
		costByDay[cost.day] += cost.value
	}
	return costByDay
}

func main() {
	costs := []cost{
		{day: 0, value: 100.0},
		{day: 1, value: 50.0},
		{day: 0, value: 25.0},
		{day: 2, value: 75.0},
		{day: 3, value: 30.0},
		{day: 4, value: 30.0},
		{day: 3, value: 30.0},
	}

	// Calculate the total cost per day
	dailyCosts := getDailyCost(costs)

	for day, totalCost := range dailyCosts {
		fmt.Printf("Day %d: %.2f\n", day+1, totalCost) // adding 1 to day for "Day 1" instead of "Day 0"
	}
}
