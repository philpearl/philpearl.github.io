package badgo

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"

	"github.com/unravelin/null"
)

func BenchmarkEncodeMarshaler(b *testing.B) {
	b.ReportAllocs()

	m := struct {
		A null.Int
		B time.Time
		C time.Time
		D null.String
	}{
		A: null.IntFrom(42),
		B: time.Now(),
		C: time.Now().Add(-time.Hour),
		D: null.StringFrom(`hello`),
	}

	b.RunParallel(func(pb *testing.PB) {
		enc := json.NewEncoder(ioutil.Discard)

		for pb.Next() {
			if err := enc.Encode(&m); err != nil {
				b.Fatal("Encode:", err)
			}
		}
	})
}
