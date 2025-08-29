package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	m1 := Constructor(matrix)

	fmt.Printf("m2: %#v\n", m1)

	// If your type has methods, you can call them here, e.g.:
	//fmt.Println(m2.SumRegion(0, 0, 1, 2))

}

type NumMatrix struct {
	matrix     [][]int
	prefixSums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	n := len(matrix)
	m := 0
	if n > 0 {
		m = len(matrix[0])
	}

	sumMatrix := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		sumMatrix[i] = make([]int, m+1)
	}

	fmt.Printf("m2: %#v\n", sumMatrix)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			sumMatrix[i+1][j+1] = matrix[i][j] + sumMatrix[i][j+1] + sumMatrix[i+1][j] - sumMatrix[i][j]
		}
	}
	fmt.Printf("m2: %#v\n", sumMatrix)

	return NumMatrix{
		matrix:     matrix,
		prefixSums: sumMatrix,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.prefixSums[row2+1][col2+1] - this.prefixSums[row1][col2+1] - this.prefixSums[row2+1][col1] + this.prefixSums[row1][col1]
}
