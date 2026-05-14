package main

import "fmt"

func printMatrix(matrix [3][3]int) {
	for k := 0; k < len(matrix); k++ {
		for l := 0; l < len(matrix[k]); l++ {
			fmt.Printf("%3v", matrix[k][l])
		}
		fmt.Println()
	}
}

func main() {
	matrix := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Printf("Matrix ist: %v\n", matrix)
	matrixNeu := [3][3]int{}

	printMatrix(matrix)

	for i := range 3 {
		for j := range matrix[i] {
			matrixNeu[j][i] = matrix[i][j]
			//fmt.Printf{"%3v", matrixNeu[i][j]}
		}
	}

	fmt.Println("Die transponierte Matrix:")
	printMatrix(matrixNeu)

}
