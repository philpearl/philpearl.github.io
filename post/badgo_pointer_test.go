package badgo

import "testing"

const bigStructSize = 1000

type bigStruct struct {
	a [bigStructSize]int
}

func newBigStruct() bigStruct {
	var b bigStruct
	for i := 0; i < bigStructSize; i++ {
		b.a[i] = i
	}
	return b
}

func newBigStructPtr() *bigStruct {
	var b bigStruct
	for i := 0; i < bigStructSize; i++ {
		b.a[i] = i
	}
	return &b
}

func BenchmarkStructReturnValue(b *testing.B) {
	b.ReportAllocs()

	t := 0
	for i := 0; i < b.N; i++ {
		v := newBigStruct()
		t += v.a[0]
	}
}

func BenchmarkStructReturnPointer(b *testing.B) {
	b.ReportAllocs()

	t := 0
	for i := 0; i < b.N; i++ {
		v := newBigStructPtr()
		t += v.a[0]
	}
}
