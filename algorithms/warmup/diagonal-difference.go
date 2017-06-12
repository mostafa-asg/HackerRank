//diagonal difference
//problem: https://www.hackerrank.com/challenges/diagonal-difference

package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	matrix := makeMatrix(n)

	fillMatrix(matrix)

	fmt.Println(calculateDiagonalDiff(matrix))

}

func calculateDiagonalDiff(matrix [][]int) int {

	primary := 0
	secondary := 0

	matLen := len(matrix)

	for i := 0; i < matLen; i++ {

		primary += matrix[i][i]
		secondary += matrix[i][matLen-i-1]

	}

	return abs(primary - secondary)
}

func abs(n int) int {

	if n < 0 {
		return -n
	}

	return n

}

func fillMatrix(matrix [][]int) {

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

}

func makeMatrix(n int) [][]int {

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	return matrix

}
