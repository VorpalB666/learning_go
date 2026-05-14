package main

import "fmt"

func main() {
	matrix := [6][5]int{{15, 8, 1, 24, 17}, {16, 14, 7, 5, 23}, {22, 20, 13, 6, 4}, {3, 21, 19, 12, 10},
		{9, 2, 25, 18, 11}, {27, 65, 13, 31, -1}}

	fehlerCount := 0

	fmt.Println(matrix)

	for j := 0; j < 5; j++ {
		var spaltenSumme int
		for i := 0; i < 5; i++ {
			fmt.Printf("i ist jetzt: %v, j ist jetzt: %v\n", i, j)
			fmt.Println(matrix[i][j])
			if i == 0 || i%2 == 0 {
				spaltenSumme = spaltenSumme + matrix[i][j]
			} else {
				spaltenSumme = spaltenSumme - matrix[i][j]
			}
		}
		fmt.Printf("Die Spaltensumme von Spalte %v ist: %v\n", j+1, spaltenSumme)
		if spaltenSumme == matrix[5][j] {
			fmt.Println("stimmt")
		} else {
			fmt.Println("stimmt nicht")
			matrix[5][j] = spaltenSumme
			fehlerCount += 1
		}
	}
	fmt.Println(matrix)
	fmt.Printf("Die Matrixe hatte %v Fehler", fehlerCount)
}
