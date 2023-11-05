package badgo

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	vals := []int{
		7, 3, 2, 293, 1, 34, 4, 99,
	}

	sorted := mergeSort(vals)
	assert.Equal(t, []int{1, 2, 3, 4, 7, 34, 99, 293}, sorted)
}

func TestSort2(t *testing.T) {
	vals := []int{
		7, 3, 2, 293, 1, 34, 4, 99,
	}

	sorted := mergeSort2(vals)
	assert.Equal(t, []int{1, 2, 3, 4, 7, 34, 99, 293}, sorted)
}

func BenchmarkSort(b *testing.B) {
	vals := rand.Perm(1000000)
	toSort := make([]int, len(vals))

	b.Run("std", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			copy(toSort, vals)
			sort.Ints(toSort)
		}
	})

	b.Run("merge", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			copy(toSort, vals)
			mergeSort(toSort)
		}
	})

	b.Run("merge2", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			copy(toSort, vals)
			mergeSort2(toSort)
		}
	})
}
