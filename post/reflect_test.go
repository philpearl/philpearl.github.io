package badgo

import "testing"

func TestPopulateReflect(t *testing.T) {
	var m SimpleStruct
	if err := populateStructReflectCache(&m); err != nil {
		t.Fatal(err)
	}
	if m.B != 42 {
		t.Fatalf("unexpected value %d for B", m.B)
	}
}

func BenchmarkPopulate(b *testing.B) {
	b.ReportAllocs()
	var m SimpleStruct
	for i := 0; i < b.N; i++ {
		populateStruct(&m)
		if m.B != 42 {
			b.Fatalf("unexpected value %d for B", m.B)
		}
	}
}

func BenchmarkPopulateReflect(b *testing.B) {
	b.ReportAllocs()
	var m SimpleStruct
	for i := 0; i < b.N; i++ {
		if err := populateStructReflect(&m); err != nil {
			b.Fatal(err)
		}
		if m.B != 42 {
			b.Fatalf("unexpected value %d for B", m.B)
		}
	}
}

func BenchmarkPopulateReflectCache(b *testing.B) {
	b.ReportAllocs()
	var m SimpleStruct
	for i := 0; i < b.N; i++ {
		if err := populateStructReflectCache(&m); err != nil {
			b.Fatal(err)
		}
		if m.B != 42 {
			b.Fatalf("unexpected value %d for B", m.B)
		}
	}
}

func BenchmarkPopulateUnsafe(b *testing.B) {
	b.ReportAllocs()
	var m SimpleStruct
	for i := 0; i < b.N; i++ {
		if err := populateStructUnsafe(&m); err != nil {
			b.Fatal(err)
		}
		if m.B != 42 {
			b.Fatalf("unexpected value %d for B", m.B)
		}
	}
}

func BenchmarkPopulateUnsafe2(b *testing.B) {
	b.ReportAllocs()
	var m SimpleStruct
	for i := 0; i < b.N; i++ {
		if err := populateStructUnsafe2(&m); err != nil {
			b.Fatal(err)
		}
		if m.B != 42 {
			b.Fatalf("unexpected value %d for B", m.B)
		}
	}
}

func BenchmarkPopulateUnsafe3(b *testing.B) {
	b.ReportAllocs()
	var m SimpleStruct

	descriptor, err := describeType((*SimpleStruct)(nil))
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		if err := populateStructUnsafe3(&m, descriptor); err != nil {
			b.Fatal(err)
		}
		if m.B != 42 {
			b.Fatalf("unexpected value %d for B", m.B)
		}
	}
}
