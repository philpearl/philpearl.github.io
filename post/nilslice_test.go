package badgo

import (
	fmt "fmt"
	"reflect"
	"testing"
	"unsafe"
)

func printSlice(s []byte) {
	sh := *(*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Printf("%#v\n", sh)
}

func BenchmarkEmptySlice(b *testing.B) {

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		t := 0
		for pb.Next() {
			s := getSlice(0)
			t += len(s)
		}
		if t == 37 {
			b.Fatal("shoudl be zero", t)
		}
	})
}

//go:noinline
func getSlice(len int) []byte {
	return nil
}
