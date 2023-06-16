package main

import (
	"fmt"
	"sync"
)

func multiplyMatrix(a, b [][]int) [][]int {
	m := len(a)
	n := len(b[0])
	result := make([][]int, m)
	for i := range result {
		result[i] = make([]int, n)
	}
	var wg sync.WaitGroup
	wg.Add(m)
	for i := 0; i < m; i++ {
		go calculate(&wg, i, a, b, result)
	}
	wg.Wait()
	return result
}

func calculate(wg *sync.WaitGroup, row int, a, b, result [][]int) {
	defer wg.Done()
	for j := 0; j < len(b[0]); j++ {
		for k := 0; k < len(b); k++ {
			result[row][j] += a[row][k] * b[k][j]
		}
	}
}

func main() {
	matrixA := [][]int{{1, 2, 3}, {4, 5, 6}}
	matrixB := [][]int{{7, 8}, {9, 10}, {11, 12}}

	result := multiplyMatrix(matrixA, matrixB)

	fmt.Println("Result Matrix:")
	for _, row := range result {
		fmt.Println(row)
	}
}
