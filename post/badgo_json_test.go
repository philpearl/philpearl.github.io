package badgo

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

type mystruct struct {
	A int       `json:"a,omitempty"`
	B string    `json:"b,omitempty"`
	C time.Time `json:"c,omitempty"`
}

func BenchmarkJSONMarshal(b *testing.B) {
	b.ReportAllocs()
	tt := time.Now()
	var data = mystruct{
		A: 42,
		B: "42",
		C: tt,
	}
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(&data)
		if err != nil {
			b.Errorf("failed to marshal json. %s", err)
		}
	}
}

func BenchmarkJSONMarshalEncoder(b *testing.B) {
	b.ReportAllocs()
	tt := time.Now()
	var data = mystruct{
		A: 42,
		B: "42",
		C: tt,
	}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	for i := 0; i < b.N; i++ {
		err := enc.Encode(&data)
		if err != nil {
			b.Errorf("failed to marshal json. %s", err)
		}
		buf.Reset()
	}
}

func BenchmarkTimeMarshal(b *testing.B) {
	b.ReportAllocs()
	var t time.Time

	for i := 0; i < b.N; i++ {
		_, err := t.MarshalJSON()
		if err != nil {
			b.Errorf("failed to marshal. %s", err)
		}
	}
}

func BenchmarkRawMarshal(b *testing.B) {
	b.ReportAllocs()
	v := struct {
		A json.RawMessage `json:"a,omitempty"`
	}{
		A: json.RawMessage(`["hello"]`),
	}

	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(v)
		if err != nil {
			b.Errorf("failed to marshal. %s", err)
		}
	}
}

func BenchmarkValueInterface(b *testing.B) {
	b.ReportAllocs()
	var s = &struct {
		T time.Time
	}{}

	sv := reflect.ValueOf(s)
	v := sv.Elem().Field(0)

	for i := 0; i < b.N; i++ {
		va := v.Addr()
		m, ok := va.Interface().(json.Marshaler)
		if !ok {
			b.Errorf("interface is nil")
		}
		_, err := m.MarshalJSON()
		if err != nil {
			b.Errorf("failed to marshal. %s", err)
		}
	}
}
