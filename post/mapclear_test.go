package badgo_test

import (
	"strconv"
	"testing"
)

func BenchmarkMapClearing(b *testing.B) {
	m := make(map[string]struct{}, 100_000)

	for range b.N {
		clear(m)
	}
}

func BenchmarkMapClearingWithItems(b *testing.B) {
	for _, size := range []int{100, 1_000, 10_000, 100_000} {
		m := make(map[string]struct{}, size)

		b.Run(strconv.Itoa(size), func(b *testing.B) {
			for range b.N {
				m["hat"] = struct{}{}
				clear(m)
			}
		})
	}
}

func BenchmarkMapAllocation(b *testing.B) {
	for _, size := range []int{100, 1_000, 10_000, 100_000} {
		b.Run(strconv.Itoa(size), func(b *testing.B) {
			for range b.N {
				m := make(map[string]struct{}, size)
				_ = m
			}
		})
	}
}
