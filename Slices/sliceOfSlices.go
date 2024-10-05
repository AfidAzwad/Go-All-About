package main

import "fmt"

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, 0)
	for i := range rows {
		row := make([]int, 0)
		for j := range cols {
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func printMatrix(rows, cols int) {
	fmt.Printf("\nMatrix rows: %d\nMatrix cols: %d\n", rows, cols)
	matrix := createMatrix(rows, cols)
	for i := range matrix {
		fmt.Println(matrix[i])
	}
}

func main() {
	rows := 10
	cols := 10

	printMatrix(rows, cols)
}
