package matrix_test

import (
	"course-info/matrix"
	"testing"
)

func BenchmarkColumnTraverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		matrix.ColumnTraverse()
	}
}

func BenchmarkRowTraverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		matrix.RowTraverse()
	}
}
