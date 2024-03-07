package main

import "fmt"

func main() {

	var matrix [3][3]int

	fmt.Println("Enter Matrix elements: \n")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("Elements: matrix[%d][%d]: ", i , j)
			fmt.Scanf("%d", &matrix[i][j])
		}
	}

	var sum1, sum2 int
	for i := 0; i < 3; i++ {
		sum1 += matrix[i][i]
		sum2 += matrix[i][2-i]
	}

	fmt.Printf("Sum of Diagonal 1: %d\n", sum1)
	fmt.Printf("Sum of Diagonal 2: %d\n", sum2)
}