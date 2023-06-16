package main

import "fmt"

func factorial(n int, arr []int) int {
	if arr[n] != 0 {
		return arr[n]
	}

	return n * factorial(n-1, arr)
}

func main() {
	num := 5

	arr := make([]int, num+1)
	arr[0] = 1

	result := factorial(num, arr)
	fmt.Printf("The factorial of %d is %d\n", num, result)
}
