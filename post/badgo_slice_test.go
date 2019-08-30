package badgo

import "testing"

type MyStruct struct {
	A int
	B int
}

func BenchmarkSlicePointers(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		slice := make([]*MyStruct, 0, 100)
		for j := 0; j < 100; j++ {
			slice = append(slice, &MyStruct{A: j, B: j + 1})
		}
	}
}

func BenchmarkSliceNoPointers(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		slice := make([]MyStruct, 0, 100)
		for j := 0; j < 100; j++ {
			slice = append(slice, MyStruct{A: j, B: j + 1})
		}
	}
}

func BenchmarkSliceHybrid(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		slice := make([]MyStruct, 0, 100)
		for j := 0; j < 100; j++ {
			slice = append(slice, MyStruct{A: j, B: j + 1})
		}

		slicep := make([]*MyStruct, len(slice))
		for j := range slice {
			slicep[j] = &slice[j]
		}
	}
}
