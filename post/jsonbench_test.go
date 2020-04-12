package badgo

import (
	"encoding/json"
	"testing"
	"time"

	jsoniter "github.com/json-iterator/go"
	easyjson "github.com/mailru/easyjson"
	"github.com/modern-go/reflect2"
	"github.com/unravelin/null"
)

func BenchmarkEncodeMarshaler(b *testing.B) {
	jsoniter.RegisterTypeEncoder(reflect2.TypeOf(null.Int{}).String(), nullIntCodec{})
	jsoniter.RegisterTypeEncoder(reflect2.TypeOf(null.String{}).String(), nullStringCodec{})
	jsoniter.RegisterTypeEncoder("time.Time", timeCodec{})

	m := myteststruct{
		A: null.IntFrom(42),
		B: time.Now(),
		C: time.Now().Add(-time.Hour),
		D: null.StringFrom(`hello`),
	}

	b.Run("encoding/json", func(b *testing.B) {
		b.ReportAllocs()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if _, err := json.Marshal(&m); err != nil {
					b.Fatal("Encode:", err)
				}
			}
		})
	})

	b.Run("jsoniter", func(b *testing.B) {
		b.ReportAllocs()
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if _, err := json.Marshal(&m); err != nil {
					b.Fatal("Encode:", err)
				}
			}
		})
	})

	b.Run("easyjson", func(b *testing.B) {
		b.ReportAllocs()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				if _, err := easyjson.Marshal(&m); err != nil {
					b.Fatal("Encode:", err)
				}
			}
		})
	})
}

func BenchmarkDecodeMarshaler(b *testing.B) {
	jsoniter.RegisterTypeDecoder(reflect2.TypeOf(null.Int{}).String(), nullIntCodec{})
	jsoniter.RegisterTypeDecoder(reflect2.TypeOf(null.String{}).String(), nullStringCodec{})
	jsoniter.RegisterTypeDecoder("time.Time", timeCodec{})

	m := myteststruct{
		A: null.IntFrom(42),
		B: time.Now(),
		C: time.Now().Add(-time.Hour),
		D: null.StringFrom(`hello`),
	}

	data, err := json.Marshal(&m)
	if err != nil {
		b.Fatal(err)
	}

	b.Run("encoding/json", func(b *testing.B) {
		b.ReportAllocs()
		b.RunParallel(func(pb *testing.PB) {
			var n myteststruct
			for pb.Next() {
				if err := json.Unmarshal(data, &n); err != nil {
					b.Fatal(err)
				}
			}
		})
	})

	b.Run("jsoniter", func(b *testing.B) {
		b.ReportAllocs()
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		b.RunParallel(func(pb *testing.PB) {
			var n myteststruct
			for pb.Next() {
				if err := json.Unmarshal(data, &n); err != nil {
					b.Fatal(err)
				}
			}
		})
	})

	b.Run("easyjson", func(b *testing.B) {
		b.ReportAllocs()
		b.RunParallel(func(pb *testing.PB) {
			var n myteststruct
			for pb.Next() {
				if err := easyjson.Unmarshal(data, &n); err != nil {
					b.Fatal(err)
				}
			}
		})
	})
}
