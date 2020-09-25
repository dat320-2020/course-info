package matrix

import "fmt"

const (
	rows = 4 * 1024
	cols = 4 * 1024
)

var matrix [rows][cols]byte

func init() {
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if row%2 == 0 {
				matrix[row][col] = 0xFF
			} else {
				matrix[row][col] = 0x11
			}
		}
	}
	fmt.Println("Elements in the matrix", rows*cols)
}

func ColumnTraverse() int {
	var cnt int
	for j := 0; j < cols; j++ {
		for i := 0; i < rows; i++ {
			if matrix[i][j] == 0xFF {
				cnt++
			}
		}
	}
	return cnt
}

func RowTraverse() int {
	var cnt int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0xFF {
				cnt++
			}
		}
	}
	return cnt
}
