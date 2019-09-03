package badgo

import (
	"fmt"
	"strconv"
	"testing"
)

func BenchmarkSprintfNumber(b *testing.B) {
	b.ReportAllocs()
	vals := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		vals[i] = fmt.Sprintf("%d", i+100)
	}
}

func BenchmarkSprintfStrconvNumber(b *testing.B) {
	b.ReportAllocs()
	vals := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		vals[i] = strconv.FormatInt(int64(i+100), 10)
	}
}

func BenchmarkBoolSprintf(b *testing.B) {
	b.ReportAllocs()
	vals := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		vals[i] = fmt.Sprintf("%t", i&1 == 0)
	}
}

func BenchmarkBoolStrconv(b *testing.B) {
	b.ReportAllocs()
	vals := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		vals[i] = strconv.FormatBool(i&1 == 0)
	}
}

func BenchmarkBoolTagSprintfAdd(b *testing.B) {
	b.ReportAllocs()
	vals := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		vals[i] = "working:" + fmt.Sprintf("%t", i&1 == 0)
	}
}

func BenchmarkBoolTagSprintf(b *testing.B) {
	b.ReportAllocs()
	vals := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		vals[i] = fmt.Sprintf("working:%t", i&1 == 0)
	}
}

func BenchmarkBoolTagIf(b *testing.B) {
	b.ReportAllocs()
	vals := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		isWorking := i&1 == 0
		if isWorking {
			vals[i] = "working:true"
		} else {
			vals[i] = "working:false"
		}
	}
}
