package badgo

import (
	"testing"
	"time"
)

func BenchmarkParseRFC3339(b *testing.B) {
	now := time.Now().UTC().Format(time.RFC3339Nano)

	for i := 0; i < b.N; i++ {
		if _, err := parseTime(now); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkParseRFC3339Compare(b *testing.B) {
	// Go has since special cased RFC3339, so to demonstrate the original
	// performance we modify the format string!
	format := time.RFC3339Nano + " hat"

	now := time.Now().UTC().Format(format)

	for i := 0; i < b.N; i++ {
		if _, err := time.Parse(format, now); err != nil {
			b.Fatal(err)
		}
	}
}
