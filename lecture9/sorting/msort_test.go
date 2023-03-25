package sorting

import "testing"

var a = []int{1, 5, 4, 4, 2, 2, 4, 66, 6345, 2, 2, 42, 4, 4, 11, 11, 1, 1, 42, 3431441, 4341, 13}

func BenchmarkMergeSortMulti(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MergeSortMulti(a)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MergeSort(a)
	}
}
