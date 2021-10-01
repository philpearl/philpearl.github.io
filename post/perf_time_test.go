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
