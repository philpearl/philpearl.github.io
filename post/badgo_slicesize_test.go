package badgo

import (
	"fmt"
	"strconv"
	"testing"
)

func numbersToStringsBad(numbers []int) []string {
	vals := []string{}
	for _, n := range numbers {
		vals = append(vals, strconv.Itoa(n))
	}
	return vals
}

func numbersToStringsBetter(numbers []int) []string {
	vals := make([]string, 0, len(numbers))
	for _, n := range numbers {
		vals = append(vals, strconv.Itoa(n))
	}
	return vals
}

func numbersToStringsBest(numbers []int) []string {
	vals := make([]string, len(numbers))
	for i, n := range numbers {
		vals[i] = strconv.Itoa(n)
	}
	return vals
}

func BenchmarkSliceConversion(b *testing.B) {
	numbers := make([]int, 100)
	for i := range numbers {
		numbers[i] = i
	}

	b.Run("bad", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			numbersToStringsBad(numbers)
		}
	})

	b.Run("better", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			numbersToStringsBetter(numbers)
		}
	})

	b.Run("best", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			numbersToStringsBest(numbers)
		}
	})
}

func numbersToStringsBadInstrumented(numbers []int) []string {
	vals := []string{}
	oldCapacity := cap(vals)
	for _, n := range numbers {
		vals = append(vals, strconv.Itoa(n))
		if capacity := cap(vals); capacity != oldCapacity {
			fmt.Printf("len(vals)=%d, cap(vals)=%d (was %d)\n", len(vals), capacity, oldCapacity)
			oldCapacity = capacity
		}
	}
	return vals
}

func TestBlah(t *testing.T) {
	numbers := make([]int, 100)
	for i := range numbers {
		numbers[i] = i
	}

	numbersToStringsBadInstrumented(numbers)
}
