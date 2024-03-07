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

	fmt.Printf("Diagonal elements of the matrix are: \n")
	for i := 0; i < 3; i++ {
		fmt.Printf("matrix[%d][%d] = %d\n", i, i, matrix[i][i])
	}
}