package badgo

import (
	"time"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/unravelin/null"
)

//easyjson:json
type myteststruct struct {
	A null.Int
	B time.Time
	C time.Time
	D null.String
}

type nullIntCodec struct{}

func (nullIntCodec) IsEmpty(ptr unsafe.Pointer) bool {
	ni := (*null.Int)(ptr)
	return !ni.Valid
}

func (nullIntCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	ni := (*null.Int)(ptr)
	if ni.Valid {
		stream.WriteInt64(ni.Int64)
	} else {
		stream.WriteNil()
	}
}

func (nullIntCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	ni := (*null.Int)(ptr)
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		ni.Valid = false
		return
	case jsoniter.NumberValue:
		ni.Valid = true
		ni.Int64 = iter.ReadInt64()
	default:
		iter.ReportError("decode null.Int", "unexpected JSON type")
	}
}

type nullStringCodec struct{}

func (nullStringCodec) IsEmpty(ptr unsafe.Pointer) bool {
	ni := (*null.String)(ptr)
	return !ni.Valid
}

func (nullStringCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	ni := (*null.String)(ptr)
	if ni.Valid {
		stream.WriteString(ni.String)
	} else {
		stream.WriteNil()
	}
}

func (nullStringCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	ni := (*null.String)(ptr)
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		ni.Valid = false
		return
	case jsoniter.StringValue:
		ni.Valid = true
		ni.String = iter.ReadString()
	default:
		iter.ReportError("decode null.String", "unexpected JSON type")
	}
}

type timeCodec struct{}

func (timeCodec) IsEmpty(ptr unsafe.Pointer) bool {
	return (*time.Time)(ptr).IsZero()
}

func (timeCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	t := (*time.Time)(ptr)

	var scratch [len(time.RFC3339Nano) + 2]byte
	b := t.AppendFormat(scratch[:0], `"`+time.RFC3339Nano+`"`)
	stream.Write(b)
}

func (timeCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	switch iter.WhatIsNext() {
	case jsoniter.NilValue:
		iter.Skip()
		*(*time.Time)(ptr) = time.Time{}
		return
	case jsoniter.StringValue:
		ts := iter.ReadStringAsSlice()

		t, err := time.Parse(time.RFC3339Nano, *(*string)(unsafe.Pointer(&ts)))
		if err != nil {
			iter.ReportError("decode time", err.Error())
			return
		}

		*(*time.Time)(ptr) = t
	default:
		iter.ReportError("decode time.Time", "unexpected JSON type")
	}
}
